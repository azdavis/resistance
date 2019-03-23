package main

import (
	"log"
	"net/http"
	"time"
)

func newServer() *http.Server {
	txLobbyMap := make(chan ToLobbyMap, 3)
	txWelcomer := make(chan *Client, 3)
	go runLobbyMap(txLobbyMap)
	go runWelcomer(txLobbyMap, txWelcomer)
	s := &http.Server{
		Handler:      NewHub(txWelcomer),
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
