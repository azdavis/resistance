package main

import (
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

// Hub is an http.Handler. It creates Clients from HTTP connections.
type Hub struct {
	up ws.Upgrader    // websocket upgrader
	tx chan<- *Client // from this to runWelcomer
}

// NewHub returns a new Hub.
func NewHub(tx chan<- *Client) *Hub {
	h := &Hub{
		up: ws.Upgrader{CheckOrigin: unsafeAllowAny},
		tx: tx,
	}
	return h
}

// unsafeAllowAny permits any *http.Request to be upgraded by a ws.Upgrader.
func unsafeAllowAny(r *http.Request) bool {
	return true
}

// ServeHTTP tries to upgrade the (w, r) pair into a websocket connection. If it
// is successful, it makes a new Client with a fresh CID and sends it along
// tx.
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	conn, err := h.up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("err ServeHTTP", err)
		return
	}
	h.tx <- NewClient(conn)
}
