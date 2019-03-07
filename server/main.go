package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("start")
	clientCh := make(chan *Client, 3)
	runLobbyMap(clientCh)
	s := &http.Server{
		Handler:      NewHub(clientCh),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
