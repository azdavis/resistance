package main

import (
	"log"
)

// PartyInfo contains info about a Party.
type PartyInfo struct {
	PID
	Leader string
}

// PartyMap represents a group of related parties. It contains one public
// fields: M, a mapping from Party IDs to Parties. Only one goroutine may
// call Add or Rm or access M at a time.
type PartyMap struct {
	M       map[PID]*Party // if M[x] = c, c.PID = x
	info    []PartyInfo    // if M[x] = c, info contains an entry for c
	nextPID PID            // PID to use for the next Add call
}

// NewPartyMap returns a new PartyMap.
func NewPartyMap() *PartyMap {
	pm := &PartyMap{
		M:       make(map[PID]*Party),
		info:    make([]PartyInfo, 0),
		nextPID: 1,
	}
	return pm
}

// Add creates and returns a Party with the given information.
func (pm *PartyMap) Add(
	leader *Client,
	send chan *Client,
	done chan PID,
) *Party {
	pid := pm.nextPID
	pm.nextPID++
	log.Println("Add", pid)
	p := NewParty(pid, leader, send, done)
	pm.M[pid] = p
	pm.setInfo()
	return p
}

// Rm removes the Party with the given PID. It returns the Party that was
// removed. A Party with the given PID must exist in the PartyMap.
func (pm *PartyMap) Rm(cid PID) *Party {
	p, ok := pm.M[cid]
	if !ok {
		panic("not present")
	}
	log.Println("Rm", cid)
	delete(pm.M, cid)
	pm.setInfo()
	return p
}

// Info gets the current info. TODO have this sorted by PID, and use that fact
// to make setInfo more efficient
func (pm *PartyMap) Info() []PartyInfo {
	return pm.info
}

// setInfo sets the current info.
func (pm *PartyMap) setInfo() {
	info := make([]PartyInfo, 0, len(pm.M))
	for pid, party := range pm.M {
		info = append(info, PartyInfo{PID: pid, Leader: party.LeaderName()})
	}
	pm.info = info
}
