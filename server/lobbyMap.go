package main

import (
	"log"
)

const lobbyMapChLen = 5

// LobbyMap contains the Clients who are not part of a Party.
type LobbyMap struct {
	rx chan *Client // incoming clients
}

// NewLobbyMap returns a new LobbyMap. It starts a goroutine which never stops.
func NewLobbyMap() *LobbyMap {
	lb := &LobbyMap{
		rx: make(chan *Client, lobbyMapChLen),
	}
	go lb.run()
	return lb
}

// run runs the LobbyMap.
func (lb *LobbyMap) run() {
	log.Println("enter LobbyMap run")
	clients := NewClientMap()
	parties := make(map[PID]*Party)
	nextPID := PID(1)
	partiesInfo := func() []PartyInfo {
		ret := make([]PartyInfo, 0, len(parties))
		for pid, party := range parties {
			// TODO race with party.run
			if party.started {
				continue
			}
			ret = append(ret, PartyInfo{pid, party.LeaderName()})
		}
		return ret
	}
	done := make(chan PID, lobbyMapChLen)
	start := make(chan struct{}, lobbyMapChLen)
	broadcastParties := func() {
		msg := PartyChoosing{Parties: partiesInfo()}
		for _, cl := range clients.M {
			if cl.name != "" {
				cl.tx <- msg
			}
		}
	}
	for {
		select {
		case cl := <-lb.rx:
			clients.Add(cl)
			if cl.name != "" {
				cl.tx <- PartyChoosing{Parties: partiesInfo()}
			}
		case <-start:
			broadcastParties()
		case pid := <-done:
			rmPartyClients := parties[pid].clients
			for cid := range rmPartyClients.M {
				clients.Add(rmPartyClients.Rm(cid))
			}
			delete(parties, pid)
			broadcastParties()
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case NameChoose:
				log.Println("NameChoose", cid, ts.Name)
				cl, ok := clients.M[cid]
				if !ok {
					continue
				}
				if !validName(ts.Name) {
					cl.tx <- NameChoosing{Valid: false}
					continue
				}
				cl.name = ts.Name
				cl.tx <- PartyChoosing{Parties: partiesInfo()}
			case PartyChoose:
				log.Println("PartyChoose", cid, ts.PID)
				party, ok := parties[ts.PID]
				if !ok {
					continue
				}
				party.rx <- clients.Rm(cid)
			case PartyCreate:
				log.Println("PartyCreate", cid)
				_, ok := clients.M[cid]
				if !ok {
					continue
				}
				parties[nextPID] = NewParty(
					nextPID,
					clients.Rm(cid),
					lb.rx,
					done,
					start,
				)
				broadcastParties()
			}
		}
	}
}
