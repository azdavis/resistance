package main

import (
	"log"
	"net/http"
	"sync"

	ws "github.com/gorilla/websocket"
)

// unsafeAllowAny permits any *http.Request to be upgraded by a ws.Upgrader.
func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = ws.Upgrader{CheckOrigin: unsafeAllowAny}

// Hub creates Clients from HTTP connections.
type Hub struct {
	mux      *sync.Mutex  // protect nextID
	nextID   CID          // next Client ID
	clientCh chan *Client // outgoing clients
}

// NewHub returns a new Hub.
func NewHub(clientCh chan *Client) *Hub {
	h := &Hub{
		mux:      &sync.Mutex{},
		nextID:   1,
		clientCh: clientCh,
	}
	return h
}

// ServeWs tries to upgrade the (w, r) pair into a websocket connection. If it
// is successful, it makes a new Client with a fresh CID and sends it along
// clientCh.
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
	h.mux.Lock()
	id := h.nextID
	h.nextID++
	h.mux.Unlock()
	h.clientCh <- NewClient(conn, id)
}
