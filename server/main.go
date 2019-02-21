package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func unsafeAllowAny(r *http.Request) bool {
	return true
}

var up = websocket.Upgrader{CheckOrigin: unsafeAllowAny}

func serveWs(w http.ResponseWriter, r *http.Request) {
	c, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	for {
		mt, m, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", m)
		err = c.WriteMessage(mt, m)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.Println("start")
	http.HandleFunc("/", serveWs)
	http.ListenAndServe(":8080", nil)
}
