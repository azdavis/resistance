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

type Hub struct {
	connCh   chan *ws.Conn
	clientCh chan *Client
}

func NewHub(clientCh chan *Client) *Hub {
	h := &Hub{
		connCh:   make(chan *ws.Conn),
		clientCh: clientCh,
	}
	go h.run()
	return h
}

func (h *Hub) run() {
	nextID := ID(1)
	for conn := range h.connCh {
		h.clientCh <- NewClient(conn, nextID)
		nextID++
	}
}

func (h *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {
	// TODO give HTTP statuses on error
	if r.URL.Path != "/" {
		return
	}
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ServeWs", err)
		return
	}
	h.connCh <- conn
}
