package main

import (
	"log"
)

// ClientMap represents a group of related clients.
//
// It contains two public fields:
//
// 1. M, a mapping from Client IDs to Clients.
//
// 2. C, a channel on which messages from all the clients stored in M are sent,
// with associated CID information attached to each ToServer (see Action).
//
// Only one goroutine may access a ClientMap at a time.
type ClientMap struct {
	C  chan Action           // messages from the Clients in M
	M  map[CID]*Client       // if M[x] = c, c.CID = x
	qs map[CID]chan struct{} // iff M[x] = c, close(qs[x]) stops piping
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C:  make(chan Action),
		M:  make(map[CID]*Client),
		qs: make(map[CID]chan struct{}),
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
	q := make(chan struct{})
	cm.M[cl.CID] = cl
	cm.qs[cl.CID] = q
	go cm.pipe(cl.CID, cl.rx, q)
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
	close(cm.qs[cid])
	delete(cm.qs, cid)
	return cl
}

// pipe pipes messages from the chan ToServer into this ClientMap's C, tagging
// each action with the CID. pipe qs when the given q channel is closed.
func (cm *ClientMap) pipe(cid CID, ch <-chan ToServer, q <-chan struct{}) {
	log.Println("enter pipe", cid)
	for {
		select {
		case <-q:
			log.Println("exit pipe", cid)
			return
		case ts := <-ch:
			cm.C <- Action{cid, ts}
		}
	}
}
