package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	ws "github.com/gorilla/websocket"
)

// Hub is an http.Handler. It creates Clients from HTTP connections.
type Hub struct {
	up ws.Upgrader   // websocket upgrader
	tx chan<- SrvMsg // from this to runServer
	fs http.Handler
}

// NewHub returns a new Hub.
func NewHub(tx chan<- SrvMsg) *Hub {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	h := &Hub{
		up: ws.Upgrader{CheckOrigin: unsafeAllowAny},
		tx: tx,
		fs: http.FileServer(http.Dir(filepath.Dir(ex))),
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
		h.fs.ServeHTTP(w, r)
		return
	}
	conn, err := h.up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("err ServeHTTP:", err)
		return
	}
	h.tx <- NewClient(conn)
}
