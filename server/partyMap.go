package main

import (
	"log"
)

// PartyMap represents a group of related parties. It contains one public
// fields: M, a mapping from Party IDs to Parties. Only one goroutine may
// call Add or Rm or access M at a time.
type PartyMap struct {
	M       map[PID]*Party // if M[x] = c, c.PID = x
	nextPID PID            // PID to use for the next Add call
}

// NewPartyMap returns a new PartyMap.
func NewPartyMap() *PartyMap {
	pm := &PartyMap{
		M:       make(map[PID]*Party),
		nextPID: 1,
	}
	return pm
}

// Add creates and returns a Party with the given information.
func (pm *PartyMap) Add(
	leader *Client,
	send chan<- *Client,
	done chan<- PID,
) *Party {
	pid := pm.nextPID
	pm.nextPID++
	log.Println("PartyMap Add", pid)
	p := NewParty(pid, leader, send, done)
	pm.M[pid] = p
	return p
}

// Rm removes the Party with the given PID. It returns the Party that was
// removed. A Party with the given PID must exist in the PartyMap.
func (pm *PartyMap) Rm(pid PID) *Party {
	p, ok := pm.M[pid]
	if !ok {
		panic("not present")
	}
	log.Println("PartyMap Rm", pid)
	delete(pm.M, pid)
	return p
}
