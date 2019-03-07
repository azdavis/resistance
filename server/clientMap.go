package main

import (
	"log"
)

// Action is a CID + ToServer.
type Action struct {
	CID
	ToServer
}

// ClientMap represents a group of related clients. It contains two public
// fields: M, a mapping from Client IDs to Clients, and C, a channel on which
// messages from all the clients stored in M are sent, with associated CID
// information attached to each ToServer (see Action). Only one goroutine may
// call Add or Rm or access M at a time.
type ClientMap struct {
	C     chan Action           // messages from the Clients in M
	M     map[CID]*Client       // if M[x] = c, c.CID = x
	quits map[CID]chan struct{} // iff M[x] = c, close(quits[x]) stop piping
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C:     make(chan Action),
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
	log.Println("ClientMap Add", cl.CID)
	quit := make(chan struct{})
	cm.M[cl.CID] = cl
	cm.quits[cl.CID] = quit
	go cm.pipe(cl.CID, cl.rx, quit)
}

// Rm removes the Client with the given CID. It stops the piping goroutine (see
// Add). It does not close the client itself. In fact, it returns the Client
// that was removed. A Client with the given CID must exist in the ClientMap.
func (cm *ClientMap) Rm(cid CID) *Client {
	cl, ok := cm.M[cid]
	if !ok {
		panic("not present")
	}
	log.Println("ClientMap Rm", cid)
	delete(cm.M, cid)
	close(cm.quits[cid])
	delete(cm.quits, cid)
	return cl
}

// pipe pipes messages from the chan ToServer into this ClientMap's C, tagging
// each action with the CID. pipe quits when the given quit channel is closed.
func (cm *ClientMap) pipe(cid CID, ch <-chan ToServer, quit <-chan struct{}) {
	log.Println("enter pipe", cid)
	for {
		select {
		case <-quit:
			log.Println("exit pipe", cid)
			return
		case ts := <-ch:
			cm.C <- Action{cid, ts}
		}
	}
}
