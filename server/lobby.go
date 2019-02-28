package main

import (
	"log"
)

type Lobby struct {
	clientCh chan *Client
}

func NewLobby() *Lobby {
	lb := &Lobby{
		clientCh: make(chan *Client),
	}
	go lb.run()
	return lb
}

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
