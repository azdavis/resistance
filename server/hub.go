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
	connCh   chan *ws.Conn
	clientCh chan *client
}

func newHub(clientCh chan *client) *hub {
	h := &hub{
		connCh:   make(chan *ws.Conn),
		clientCh: make(chan *client),
	}
	go h.run()
	return h
}

func (h *hub) run() {
	nextID := ID(1)
	for conn := range h.connCh {
		h.clientCh <- newClient(conn, nextID)
		nextID++
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
