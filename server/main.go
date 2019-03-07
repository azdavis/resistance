package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("start")
	toLobbyMap := make(chan *Client, 3)
	toNameChooser := make(chan *Client, 3)
	go runLobbyMap(toLobbyMap)
	go runNameChooser(toLobbyMap, toNameChooser)
	s := &http.Server{
		Handler:      NewHub(toNameChooser),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
