package main

import (
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

// unsafeAllowAny permits any *http.Request to be upgraded by a ws.Upgrader.
func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = ws.Upgrader{CheckOrigin: unsafeAllowAny}

// Hub creates Clients from HTTP connections.
type Hub struct {
	connCh   chan *ws.Conn // incoming websocket connections
	clientCh chan *Client  // outgoing clients
}

// NewHub returns a new Hub. It starts a goroutine which never exits.
func NewHub(clientCh chan *Client) *Hub {
	h := &Hub{
		connCh:   make(chan *ws.Conn),
		clientCh: clientCh,
	}
	go h.run()
	return h
}

// run runs the Hub. Whenever a conn arrives on connCh, it makes a new Client
// with a fresh ID and sends it along clientCh.
func (h *Hub) run() {
	nextID := CID(1)
	for conn := range h.connCh {
		h.clientCh <- NewClient(conn, nextID)
		nextID++
	}
}

// ServeWs tries to upgrade the (w, r) pair into a websocket connection. If it
// does, it sends the connection to run.
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
