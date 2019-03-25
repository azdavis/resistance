package main

import (
	"log"
)

func main() {
	log.Println("start")
	log.Fatal(NewServer().HTTPServer.ListenAndServe())
}
