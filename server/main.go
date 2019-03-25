package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("start")
	s := NewServer()
	hs := &http.Server{
		Handler:      NewHub(s.C),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Fatal(hs.ListenAndServe())
}
