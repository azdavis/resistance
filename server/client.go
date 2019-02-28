package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

type client struct {
	id    ID
	room  ID     // if 0, no room
	name  string // if "", no name
	isSpy bool
	send  chan State
	recv  chan Action
	conn  *ws.Conn
}

func newClient(conn *ws.Conn, id ID) *client {
	cl := &client{
		id:    id,
		room:  0,
		name:  "",
		isSpy: false,
		send:  make(chan State),
		recv:  make(chan Action),
		conn:  conn,
	}
	go cl.recvFromConn()
	go cl.sendToConn()
	return cl
}

func (cl *client) recvFromConn() {
	for {
		mt, bs, err := cl.conn.ReadMessage()
		if err != nil {
			log.Println("recvFromConn", cl.id, err)
			// no further actions will be sent on recv. however, do not close recv,
			// since we may send garbage actions to listeners. only send the Close
			// Action.
			cl.recv <- Close{}
			cl.conn.Close()
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

// cl.send will be closed by the managing room when this client has been closed.
func (cl *client) sendToConn() {
	for m := range cl.send {
		err := cl.conn.WriteJSON(m)
		if err != nil {
			log.Println("sendTo", cl.id, err)
		}
	}
}
