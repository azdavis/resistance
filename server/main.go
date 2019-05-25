package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func main() {
	hs := &http.Server{
		Handler:      NewHub(NewServer().C),
		Addr:         ":" + getPort(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	log.Println("err main:", hs.ListenAndServe())
}
