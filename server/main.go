package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("start")
	lb := NewLobby()
	s := &http.Server{
		Handler:      NewHub(lb.rx),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
