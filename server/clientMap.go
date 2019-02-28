package main

import (
	"log"
)

// IDAction is an Action from a particular client with a given ID.
type IDAction struct {
	ID
	Action
}

// ClientMap represents a group of related clients. It contains two public
// fields: M, a mapping from Client IDs to Clients, and C, a channel on which
// messages from all the clients stored in M are sent, with associated ID
// information attached to each Action (see IDAction). Only one goroutine may
// call Add or Rm or access M at a time.
type ClientMap struct {
	C     chan IDAction        // messages from the Clients in M, tagged with ID
	M     map[ID]*Client       // if M[x] = c, c.id = x
	quits map[ID]chan struct{} // iff M[x] = c, close(quits[x]) will stop piping
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C:     make(chan IDAction),
		M:     make(map[ID]*Client),
		quits: make(map[ID]chan struct{}),
	}
	return cm
}

// Add adds the given Client to the map. It starts a goroutine to pipe messages
// from the given Client to this ClientMap's C. Another Client with the same ID
// must not exist in this ClientMap.
func (cm *ClientMap) Add(cl *Client) {
	_, ok := cm.quits[cl.id]
	if ok {
		panic("already present")
	}
	log.Println("Add", cl.id)
	quit := make(chan struct{})
	cm.M[cl.id] = cl
	cm.quits[cl.id] = quit
	go cm.pipe(cl.id, cl.recv, quit)
}

// Rm removes the Client with the given ID. It stops the piping goroutine (see
// Add). A Client with the given ID must exist in the ClientMap.
func (cm *ClientMap) Rm(id ID) {
	_, ok := cm.quits[id]
	if !ok {
		panic("not present")
	}
	log.Println("Rm", id)
	close(cm.M[id].send)
	close(cm.quits[id])
	delete(cm.M, id)
	delete(cm.quits, id)
}

// pipe pipes messages from the chan Action into this ClientMap's C, tagging
// each action with the id. pipe quits when the given quit channel is closed.
func (cm *ClientMap) pipe(id ID, ch chan Action, quit chan struct{}) {
	defer log.Println("exit pipe", id)
	for {
		select {
		case <-quit:
			return
		case ac := <-ch:
			cm.C <- IDAction{id, ac}
		}
	}
}
