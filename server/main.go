package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("start")
	hs := &http.Server{
		Handler:      NewHub(NewServer().C),
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	fmt.Println("err main:", hs.ListenAndServe())
}
