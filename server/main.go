package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var env string

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func main() {
	port := getPort()
	log.Println("starting on port", port)
	hs := &http.Server{
		Handler:      NewHub(NewServer().C),
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Println("err main:", hs.ListenAndServe())
}
