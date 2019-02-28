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
	M     map[CID]*Client       // if M[x] = c, c.id = x
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
	_, ok := cm.M[cl.id]
	if ok {
		panic("already present")
	}
	log.Println("Add", cl.id)
	quit := make(chan struct{})
	cm.M[cl.id] = cl
	cm.quits[cl.id] = quit
	go cm.pipe(cl.id, cl.recv, quit)
}

// Rm removes the Client with the given CID. It stops the piping goroutine (see
// Add). A Client with the given CID must exist in the ClientMap.
func (cm *ClientMap) Rm(id CID) {
	_, ok := cm.M[id]
	if !ok {
		panic("not present")
	}
	log.Println("Rm", id)
	cm.M[id].Close()
	delete(cm.M, id)
	close(cm.quits[id])
	delete(cm.quits, id)
}

// pipe pipes messages from the chan Action into this ClientMap's C, tagging
// each action with the id. pipe quits when the given quit channel is closed.
func (cm *ClientMap) pipe(id CID, ch chan Action, quit chan struct{}) {
	defer log.Println("exit pipe", id)
	for {
		select {
		case <-quit:
			return
		case ac := <-ch:
			cm.C <- CIDAction{id, ac}
		}
	}
}
