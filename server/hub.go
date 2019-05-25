package main

import (
	"log"
	"net/http"
	"net/url"

	ws "github.com/gorilla/websocket"
)

const allowedHost = "azdavis.xyz"

// Hub turns HTTP connections into Clients.
type Hub struct {
	up ws.Upgrader   // websocket upgrader
	tx chan<- SrvMsg // from this to runServer
}

// NewHub returns a new Hub.
func NewHub(tx chan<- SrvMsg) *Hub {
	return &Hub{ws.Upgrader{CheckOrigin: unsafeDebugCheckOrigin}, tx}
}

// unsafeDebugCheckOrigin returns true.
func unsafeDebugCheckOrigin(r *http.Request) bool {
	return true
}

// checkOrigin returns whether r has an Origin header which contains a
// valid URL with azdavis.xyz as its host.
func checkOrigin(r *http.Request) bool {
	origin := r.Header["Origin"]
	if len(origin) == 0 {
		return false
	}
	u, err := url.Parse(origin[0])
	if err != nil {
		return false
	}
	return u.Host == allowedHost
}

// ServeHTTP tries to upgrade the (w, r) pair into a websocket connection. If it
// is successful, it makes a new Client with a fresh CID and sends it along
// tx.
func (h *Hub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	conn, err := h.up.Upgrade(w, r, nil)
	if err != nil {
		log.Println("err ServeHTTP:", err)
		return
	}
	h.tx <- NewClient(conn)
}
