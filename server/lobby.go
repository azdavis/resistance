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
	clients := NewClientMap()
	for {
		select {
		case cl := <-lb.clientCh:
			clients.Add(cl)
		case idAc := <-clients.C:
			id := idAc.CID
			switch ac := idAc.Action.(type) {
			case Close:
				clients.Rm(id)
			case NameChoose:
				log.Println("NameChoose", id, ac.Name)
				clients.M[id].name = ac.Name
				clients.M[id].send <- PartyChoosing{
					Name:    ac.Name,
					Parties: []string{},
				}
			}
		}
	}
}
