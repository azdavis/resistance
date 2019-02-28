package main

import (
	"log"
	"net/http"
	"sync"

	ws "github.com/gorilla/websocket"
)

// Hub creates Clients from HTTP connections.
type Hub struct {
	mux     *sync.Mutex    // protect nextCID
	nextCID CID            // the next Client will have this CID
	send    chan<- *Client // outgoing clients
}

// NewHub returns a new Hub.
func NewHub(send chan<- *Client) *Hub {
	h := &Hub{
		mux:     &sync.Mutex{},
		nextCID: 1,
		send:    send,
	}
	return h
}

// unsafeAllowAny permits any *http.Request to be upgraded by a ws.Upgrader.
func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = ws.Upgrader{CheckOrigin: unsafeAllowAny}

// ServeHTTP tries to upgrade the (w, r) pair into a websocket connection. If it
// is successful, it makes a new Client with a fresh CID and sends it along
// send.
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ws" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ServeHTTP", err)
		return
	}
	h.mux.Lock()
	cid := h.nextCID
	h.nextCID++
	h.mux.Unlock()
	h.send <- NewClient(conn, cid)
}
