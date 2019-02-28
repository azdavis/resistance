package main

import (
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = ws.Upgrader{CheckOrigin: unsafeAllowAny}

type hub struct {
	conns chan *ws.Conn
}

func newHub() *hub {
	h := &hub{
		conns: make(chan *ws.Conn),
	}
	go h.run()
	return h
}

func (h *hub) run() {
	nextID := ID(1)
	clients := make(map[ID]*client)
	bc := newBigChan()
	for {
		select {
		case conn := <-h.conns:
			id := nextID
			nextID++
			client := newClient(conn, id)
			clients[id] = client
			bc.add(id, client.recv)
		case idAc := <-bc.c:
			id := idAc.ID
			switch ac := idAc.Action.(type) {
			case Close:
				bc.rm(id)
			case NameChoose:
				clients[id].name = ac.Name
				clients[id].send <- PartyChoosing{
					Name:    ac.Name,
					Parties: []string{},
				}
			}
		}
	}
}

func (h *hub) serveWs(w http.ResponseWriter, r *http.Request) {
	// TODO give HTTP statuses on error
	if r.URL.Path != "/" {
		return
	}
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("serveWs", err)
		return
	}
	h.conns <- conn
}
