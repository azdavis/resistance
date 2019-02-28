package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

type ID uint64

type Client struct {
	id    ID
	room  ID     // if 0, no room
	name  string // if "", no name
	isSpy bool
	send  chan State
	recv  chan Action
}

func NewClient(conn *ws.Conn, id ID) *Client {
	cl := &Client{
		id:    id,
		room:  0,
		name:  "",
		isSpy: false,
		send:  make(chan State),
		recv:  make(chan Action),
	}
	go cl.recvFromConn(conn)
	go cl.sendToConn(conn)
	return cl
}

func (cl *Client) recvFromConn(conn *ws.Conn) {
	defer log.Println("exit recvFromConn", cl.id)
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("recvFromConn", cl.id, err)
			// no further actions will be sent on recv. however, do not close recv,
			// since we may send garbage actions to listeners. only send the Close
			// Action.
			cl.recv <- Close{}
			conn.Close()
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		ac, err := JSONToAction(bs)
		if err != nil {
			continue
		}
		cl.recv <- ac
	}
}

// cl.send will be closed by the managing room when this Client has been closed.
func (cl *Client) sendToConn(conn *ws.Conn) {
	defer log.Println("exit sendToConn", cl.id)
	for m := range cl.send {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("sendToConn", cl.id, err)
		}
	}
}
