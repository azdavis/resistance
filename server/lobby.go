package main

import (
	"log"
)

type lobby struct {
	clientCh chan *client
}

func newLobby() *lobby {
	lb := &lobby{
		clientCh: make(chan *client),
	}
	go lb.run()
	return lb
}

func (lb *lobby) run() {
	cm := newClientMap()
	for {
		select {
		case cl := <-lb.clientCh:
			cm.add(cl)
		case idAc := <-cm.C:
			id := idAc.ID
			switch ac := idAc.Action.(type) {
			case Close:
				cm.rm(id)
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
