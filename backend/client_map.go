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
// with associated CID information attached.
//
// Only one goroutine may access a ClientMap at a time.
type ClientMap struct {
	C chan CIDToServer // messages from the Clients in M
	M map[CID]Client   // clients
}

// NewClientMap returns a new ClientMap.
func NewClientMap() *ClientMap {
	return &ClientMap{make(chan CIDToServer), make(map[CID]Client)}
}

// Add adds the given Client to the map. Another Client with the same CID
// must not exist in this ClientMap.
func (cm *ClientMap) Add(cid CID, cl Client) {
	cm.AddNoSend(cid, cl)
	cl.RecvTo(Dest{cid, cm.C})
}

// AddNoSend adds the given Client to the map, but does not direct it to begin
// sending on C. Another Client with the same CID must not exist in this
// ClientMap.
func (cm *ClientMap) AddNoSend(cid CID, cl Client) {
	_, ok := cm.M[cid]
	if ok {
		panic("already present")
	}
	cm.M[cid] = cl
}

// Rm removes the Client with the given CID. It does not close the client
// itself. In fact, it returns the Client that was removed. A Client with the
// given CID must exist in the ClientMap.
func (cm *ClientMap) Rm(cid CID) Client {
	cl, ok := cm.M[cid]
	if !ok {
		panic("not present")
	}
	delete(cm.M, cid)
	cl.RecvTo(NullDest)
	return cl
}

// ToList converts the ClientMap to a list. The order is sorted in
// increasing-CID order.
func (cm *ClientMap) ToList(names map[CID]string) []ClientInfo {
	ret := make([]ClientInfo, 0, len(cm.M))
	for cid := range cm.M {
		ret = append(ret, ClientInfo{cid, names[cid]})
	}
	sort.Slice(ret, func(i, j int) bool { return ret[i].CID < ret[j].CID })
	return ret
}

// DisconnectAll turns off the C.
func (cm *ClientMap) DisconnectAll() {
	for _, cl := range cm.M {
		cl.RecvTo(NullDest)
	}
	close(cm.C)
}

// CloseAll kills all clients in M and empties M.
func (cm *ClientMap) CloseAll() {
	for cid := range cm.M {
		cm.Rm(cid).Close()
	}
}
