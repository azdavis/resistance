package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	lb := NewLobby()
	h := NewHub(lb.clientCh)
	http.HandleFunc("/", h.ServeWs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
