package main

import (
	"sort"
)

// ClientMap represents a group of related clients.
//
// It contains two public fields:
//
// 1. M, a mapping from Client IDs to Clients.
//
// 2. C, a channel on which messages from all the clients stored in M are sent,
// with associated CID information attached (see Action).
//
// Only one goroutine may access a ClientMap at a time.
type ClientMap struct {
	C chan Action     // messages from the Clients in M
	M map[CID]*Client // if M[x] = c, c.CID = x
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	cm := &ClientMap{
		C: make(chan Action),
		M: make(map[CID]*Client),
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
	cm.M[cl.CID] = cl
	cl.SendOn(cm.C)
}

// Rm removes the Client with the given CID. It stops the piping goroutine (see
// Add). It does not close the client itself. In fact, it returns the Client
// that was removed. A Client with the given CID must exist in the ClientMap.
func (cm *ClientMap) Rm(cid CID) *Client {
	cl, ok := cm.M[cid]
	if !ok {
		panic("not present")
	}
	delete(cm.M, cid)
	cl.SendOn(SendBlocks)
	return cl
}

// ToList converts the ClientMap to a list. The order is sorted in
// increasing-CID order.
func (cm *ClientMap) ToList() []*Client {
	ret := make([]*Client, 0, len(cm.M))
	for _, cl := range cm.M {
		ret = append(ret, cl)
	}
	// TODO bad perf
	sort.Slice(ret, func(i, j int) bool { return ret[i].CID < ret[j].CID })
	return ret
}

// Clear removes all clients from M. It returns ToList (which was calculated
// before removing all the clients).
func (cm *ClientMap) Clear() []*Client {
	ret := cm.ToList()
	for cid := range cm.M {
		cm.Rm(cid)
	}
	return ret
}

// CloseAll kills all clients in M and empties M.
func (cm *ClientMap) CloseAll() {
	for cid := range cm.M {
		cm.Rm(cid).Close()
	}
}
