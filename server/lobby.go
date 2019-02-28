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

func (h *lobby) run() {
	clients := make(map[ID]*client)
	bc := newBigChan()
	addClient := func(cl *client) {
		log.Println("addClient", cl.id)
		clients[cl.id] = cl
		bc.add(cl.id, cl.recv)
	}
	rmClient := func(id ID) {
		log.Println("rmClient", id)
		delete(clients, id)
		bc.rm(id)
	}
	for {
		select {
		case cl := <-h.clientCh:
			addClient(cl)
		case idAc := <-bc.c:
			id := idAc.ID
			switch ac := idAc.Action.(type) {
			case Close:
				rmClient(id)
			case NameChoose:
				log.Println("NameChoose", id, ac.Name)
				clients[id].name = ac.Name
				clients[id].send <- PartyChoosing{
					Name:    ac.Name,
					Parties: []string{},
				}
			}
		}
	}
}
