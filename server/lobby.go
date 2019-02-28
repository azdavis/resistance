package main

import (
	"log"
)

// Lobby contains the Clients who are not part of a Party.
type Lobby struct {
	clientCh chan *Client // incoming clients
}

// NewLobby returns a new Lobby. It starts a goroutine which never exits.
func NewLobby() *Lobby {
	lb := &Lobby{
		clientCh: make(chan *Client),
	}
	go lb.run()
	return lb
}

// run runs the Lobby.
func (lb *Lobby) run() {
	cm := NewClientMap()
	for {
		select {
		case cl := <-lb.clientCh:
			cm.Add(cl)
		case idAc := <-cm.C:
			id := idAc.ID
			switch ac := idAc.Action.(type) {
			case Close:
				cm.Rm(id)
			case NameChoose:
				log.Println("NameChoose", id, ac.Name)
				cm.M[id].name = ac.Name
				cm.M[id].send <- PartyChoosing{
					Name:    ac.Name,
					Parties: []string{},
				}
			}
		}
	}
}
