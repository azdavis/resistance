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
	c := &client{
		id:    id,
		room:  0,
		name:  "",
		isSpy: false,
		send:  make(chan State),
		recv:  make(chan Action),
		conn:  conn,
	}
	go c.recvFromConn()
	go c.sendToConn()
	return c
}

func (c *client) recvFromConn() {
	for {
		mt, bs, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("recvFromConn", c.id, err)
			c.recv <- Close{}
			// no further actions will be sent on recv.
			close(c.recv)
			c.conn.Close()
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		ac, err := JSONToAction(bs)
		if err != nil {
			continue
		}
		c.recv <- ac
	}
}

// c.send will be closed by the managing room when this client has been closed.
func (c *client) sendToConn() {
	for m := range c.send {
		err := c.conn.WriteJSON(m)
		if err != nil {
			log.Println("sendTo", c.id, err)
		}
	}
}
