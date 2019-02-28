package main

import (
	"log"
)

// Lobby contains the Clients who are not part of a Party.
type Lobby struct {
	recv chan *Client // incoming clients
	done chan PID     // incoming done parties
}

// NewLobby returns a new Lobby. It starts a goroutine which never exits.
func NewLobby() *Lobby {
	const chLen = 5
	lb := &Lobby{
		recv: make(chan *Client, chLen),
		done: make(chan PID, chLen),
	}
	go lb.run()
	return lb
}

// run runs the Lobby.
func (lb *Lobby) run() {
	clients := NewClientMap()
	nextPID := PID(1)
	parties := make(map[PID]*Party)
	getPartyInfo := func() []PartyInfo {
		ret := make([]PartyInfo, 0, len(parties))
		for pid, party := range parties {
			ret = append(ret, PartyInfo{PID: pid, Name: party.name})
		}
		return ret
	}
	for {
		select {
		case cl := <-lb.recv:
			clients.Add(cl)
		case pid := <-lb.done:
			cm := parties[pid].clients
			for cid := range cm.M {
				clients.Add(cm.Rm(cid))
			}
			delete(parties, pid)
		case cidAc := <-clients.C:
			id := cidAc.CID
			switch ac := cidAc.Action.(type) {
			case Close:
				clients.Rm(id).Close()
			case NameChoose:
				log.Println("NameChoose", id, ac.Name)
				clients.M[id].name = ac.Name
				clients.M[id].send <- PartyChoosing{
					Name:    ac.Name,
					Parties: getPartyInfo(),
				}
			case PartyChoose:
				log.Println("PartyChoose", id, ac.PID)
				parties[ac.PID].recv <- clients.Rm(id)
			case PartyCreate:
				log.Println("PartyCreate", id, ac.Name)
				parties[nextPID] = NewParty(nextPID, ac.Name, clients.Rm(id), lb.done)
				nextPID++
			}
		}
	}
}
