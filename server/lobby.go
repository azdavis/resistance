package main

import (
	"log"
)

const lobbyChLen = 5

// Lobby contains the Clients who are not part of a Party.
type Lobby struct {
	recv chan *Client // incoming clients
}

// NewLobby returns a new Lobby. It starts a goroutine which never exits.
func NewLobby() *Lobby {
	lb := &Lobby{
		recv: make(chan *Client, lobbyChLen),
	}
	go lb.run()
	return lb
}

// run runs the Lobby.
func (lb *Lobby) run() {
	log.Println("enter Lobby run")
	clients := NewClientMap()
	parties := NewPartyMap()
	partiesInfo := func() []PartyInfo {
		ret := make([]PartyInfo, 0, len(parties.M))
		for pid, party := range parties.M {
			ret = append(ret, PartyInfo{pid, party.LeaderName()})
		}
		return ret
	}
	done := make(chan PID, lobbyChLen)
	broadcastParties := func() {
		msg := PartyChoosing{Parties: partiesInfo()}
		for _, cl := range clients.M {
			if cl.name != "" {
				cl.send <- msg
			}
		}
	}
	for {
		select {
		case cl := <-lb.recv:
			clients.Add(cl)
			if cl.name != "" {
				cl.send <- PartyChoosing{Parties: partiesInfo()}
			}
		case pid := <-done:
			rmPartyClients := parties.Rm(pid).clients
			for cid := range rmPartyClients.M {
				clients.Add(rmPartyClients.Rm(cid))
			}
			broadcastParties()
		case ac := <-clients.C:
			cid := ac.CID
			switch ac := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case NameChoose:
				log.Println("NameChoose", cid, ac.Name)
				cl, ok := clients.M[cid]
				if !ok {
					continue
				}
				if !validName(ac.Name) {
					cl.send <- NameChoosing{Valid: false}
					continue
				}
				cl.name = ac.Name
				cl.send <- PartyChoosing{Parties: partiesInfo()}
			case PartyChoose:
				log.Println("PartyChoose", cid, ac.PID)
				party, ok := parties.M[ac.PID]
				if !ok {
					continue
				}
				party.recv <- clients.Rm(cid)
			case PartyCreate:
				log.Println("PartyCreate", cid)
				_, ok := clients.M[cid]
				if !ok {
					continue
				}
				parties.Add(clients.Rm(cid), lb.recv, done)
				broadcastParties()
			}
		}
	}
}
