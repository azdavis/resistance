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
	actionCh chan IDAction
	closeCh  chan ID
	connCh   chan *ws.Conn
}

func newHub() *hub {
	h := &hub{
		actionCh: make(chan IDAction),
		closeCh:  make(chan ID),
		connCh:   make(chan *ws.Conn),
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

func (h *hub) recvFrom(conn *ws.Conn, id ID) {
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("recvFrom", id, err)
			h.closeCh <- id
			conn.Close()
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		ac, err := JSONToAction(bs)
		if err != nil {
			continue
		}
		h.actionCh <- IDAction{id, ac}
	}
}

func (h *hub) sendTo(conn *ws.Conn, id ID, ms chan State) {
	for m := range ms {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("sendTo", id, err)
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
	h.connCh <- conn
}
