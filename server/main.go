package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")
	h := newHub()
	go h.run()
	http.HandleFunc("/", h.serveWs)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
