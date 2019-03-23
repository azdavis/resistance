package main

import (
	"log"
	"net/http"
	"time"
)

func newServer() *http.Server {
	toLobbyMap := make(chan *Client, 3)
	toWelcomeLobby := make(chan *Client, 3)
	go runLobbyMap(toLobbyMap)
	go runWelcomeLobby(toLobbyMap, toWelcomeLobby)
	s := &http.Server{
		Handler:      NewHub(toWelcomeLobby),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return s
}

func main() {
	log.Println("start")
	log.Fatal(newServer().ListenAndServe())
}
