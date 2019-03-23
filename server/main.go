package main

import (
	"log"
	"net/http"
	"time"
)

func newServer() *http.Server {
	toLobbyMap := make(chan *Client, 3)
	toWelcomer := make(chan *Client, 3)
	go runLobbyMap(toLobbyMap)
	go runWelcomer(toLobbyMap, toWelcomer)
	s := &http.Server{
		Handler:      NewHub(toWelcomer),
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
