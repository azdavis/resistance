package main

import (
	"log"
)

// CIDAction is an Action from a particular client with a given CID.
type CIDAction struct {
	CID
	Action
}

// ClientMap represents a group of related clients. It contains two public
// fields: M, a mapping from Client IDs to Clients, and C, a channel on which
// messages from all the clients stored in M are sent, with associated CID
// information attached to each Action (see CIDAction). Only one goroutine may
// call Add or Rm or access M at a time.
type ClientMap struct {
	C     chan CIDAction        // messages from the Clients in M, tagged with CID
	M     map[CID]*Client       // if M[x] = c, c.CID = x
	quits map[CID]chan struct{} // iff M[x] = c, close(quits[x]) will stop piping
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C:     make(chan CIDAction),
		M:     make(map[CID]*Client),
		quits: make(map[CID]chan struct{}),
	}
	return cm
}

// Add adds the given Client to the map. It starts a goroutine to pipe messages
// from the given Client to this ClientMap's C. Another Client with the same CID
// must not exist in this ClientMap.
func (cm *ClientMap) Add(cl *Client) {
	_, ok := cm.M[cl.CID]
	if ok {
		panic("already present")
	}
	log.Println("Add", cl.CID)
	quit := make(chan struct{})
	cm.M[cl.CID] = cl
	cm.quits[cl.CID] = quit
	go cm.pipe(cl.CID, cl.recv, quit)
}

// Rm removes the Client with the given CID. It stops the piping goroutine (see
// Add). It does not close the client itself. In fact, it returns the client
// that was removed. A Client with the given CID must exist in the ClientMap.
func (cm *ClientMap) Rm(cid CID) *Client {
	cl, ok := cm.M[cid]
	if !ok {
		panic("not present")
	}
	log.Println("Rm", cid)
	delete(cm.M, cid)
	close(cm.quits[cid])
	delete(cm.quits, cid)
	return cl
}

// pipe pipes messages from the chan Action into this ClientMap's C, tagging
// each action with the CID. pipe quits when the given quit channel is closed.
func (cm *ClientMap) pipe(CID CID, ch chan Action, quit chan struct{}) {
	for {
		select {
		case <-quit:
			log.Println("exit pipe", CID)
			return
		case ac := <-ch:
			cm.C <- CIDAction{CID, ac}
		}
	}
}
