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
			ret = append(ret, PartyInfo{PID: pid, Leader: party.LeaderName()})
		}
		return ret
	}
	partyInfo := getPartyInfo()
	for {
		select {
		case cl := <-lb.recv:
			clients.Add(cl)
		case pid := <-lb.done:
			cm := parties[pid].clients
			delete(parties, pid)
			partyInfo = getPartyInfo()
			for cid := range cm.M {
				cl := cm.Rm(cid)
				cl.send <- PartyDisbanded{Parties: partyInfo}
				clients.Add(cl)
			}
		case ac := <-clients.C:
			cid := ac.CID
			switch ac := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case NameChoose:
				log.Println("NameChoose", cid, ac.Name)
				clients.M[cid].name = ac.Name
				clients.M[cid].send <- PartyChoosing{
					Parties: partyInfo,
				}
			case PartyChoose:
				log.Println("PartyChoose", cid, ac.PID)
				parties[ac.PID].recv <- clients.Rm(cid)
			case PartyCreate:
				log.Println("PartyCreate", cid)
				parties[nextPID] = NewParty(
					nextPID,
					clients.Rm(cid),
					lb.recv,
					lb.done,
				)
				nextPID++
				partyInfo = getPartyInfo()
				for _, cl := range clients.M {
					cl.send <- PartyChoosing{
						Parties: partyInfo,
					}
				}
			}
		}
	}
}
