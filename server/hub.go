package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = websocket.Upgrader{CheckOrigin: unsafeAllowAny}

type hub struct {
	actionCh chan Action
	closeCh  chan ID
	connCh   chan *websocket.Conn
}

func newHub() *hub {
	h := &hub{
		actionCh: make(chan Action),
		closeCh:  make(chan ID),
		connCh:   make(chan *websocket.Conn),
	}
	go h.run()
	return h
}

func (h *hub) run() {
	conns := make(map[ID]chan State)
	nextID := ID(1)
	for {
		select {
		case msg := <-h.actionCh:
			log.Println(msg)
		case id := <-h.closeCh:
			ms, ok := conns[id]
			if ok {
				close(ms)
				delete(conns, id)
			}
		case conn := <-h.connCh:
			ms := make(chan State)
			conns[nextID] = ms
			go h.recvFrom(conn, nextID)
			go h.sendTo(conn, nextID, ms)
			nextID++
		}
	}
}

func (h *hub) recvFrom(conn *websocket.Conn, id ID) {
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("recvFrom", id, err)
			h.closeCh <- id
			conn.Close()
			return
		}
		if mt != websocket.TextMessage {
			continue
		}
		ac, err := JSONToAction(bs)
		if err != nil {
			continue
		}
		h.actionCh <- ac
	}
}

func (h *hub) sendTo(conn *websocket.Conn, id ID, ms chan State) {
	for m := range ms {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("sendTo", id, err)
		}
	}
}

func (h *hub) serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("serveWs", err)
		return
	}
	h.connCh <- conn
}
