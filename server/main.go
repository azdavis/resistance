package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	lb := newLobby()
	h := newHub(lb.clientCh)
	http.HandleFunc("/", h.serveWs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
