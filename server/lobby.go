package main

import (
	"log"
)

// Lobby contains the Clients who are not part of a Party.
type Lobby struct {
	recv chan *Client // incoming clients
}

// NewLobby returns a new Lobby. It starts a goroutine which never exits.
func NewLobby() *Lobby {
	const chLen = 5
	lb := &Lobby{
		recv: make(chan *Client, chLen),
	}
	go lb.run()
	return lb
}

// run runs the Lobby.
func (lb *Lobby) run() {
	clients := NewClientMap()
	for {
		select {
		case cl := <-lb.recv:
			clients.Add(cl)
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
					Parties: []PartyInfo{},
				}
			}
		}
	}
}
