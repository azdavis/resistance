package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("start")
	s := NewServer()
	hs := &http.Server{
		Handler:      NewHub(s.C),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	fmt.Println("err main:", hs.ListenAndServe())
}
