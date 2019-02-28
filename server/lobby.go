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
	cm := newClientMap()
	addClient := func(cl *client) {
		log.Println("addClient", cl.id)
		clients[cl.id] = cl
		cm.add(cl)
	}
	rmClient := func(id ID) {
		log.Println("rmClient", id)
		delete(clients, id)
		cm.rm(id)
	}
	for {
		select {
		case cl := <-h.clientCh:
			addClient(cl)
		case idAc := <-cm.c:
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
