package main

import (
	"fmt"
	"time"

	ws "github.com/gorilla/websocket"
)

// NullDest is a Dest which will never allow sending.
var NullDest = Dest{0, make(chan<- Action)}

// PingPeriod is the frequency with which pings are sent.
const PingPeriod = 40 * time.Second

// Client is a player of the game.
type Client struct {
	tx      chan ToClient // orders for the client
	rx      chan ToServer // requests from the client
	newDest chan Dest     // what to update the ultimate destination of rx to
	q       chan struct{} // close on Close
	conn    *ws.Conn      // the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection, if it wasn't nil.
func NewClient(conn *ws.Conn) Client {
	cl := Client{
		tx:      make(chan ToClient, 3),
		rx:      make(chan ToServer, 3),
		newDest: make(chan Dest),
		q:       make(chan struct{}),
		conn:    conn,
	}
	go cl.manageDest()
	if conn != nil {
		go cl.readFromConn()
		go cl.writeToConn()
	}
	return cl
}

// Close quits all goroutines started with NewClient. It should be called
// exactly once. Usually this is called after receiving a Close{} on rx.
func (cl Client) Close() {
	if cl.conn != nil {
		cl.conn.Close()
	}
	close(cl.q)
	close(cl.tx)
}

// RecvTo updates dest. It returns only once dest has been updated.
func (cl Client) RecvTo(dest Dest) {
	cl.newDest <- dest
}

func (cl Client) manageDest() {
	dest := NullDest
	var m ToServer
recv:
	for {
		select {
		case <-cl.q:
			return
		case dest = <-cl.newDest:
		case m = <-cl.rx:
			goto send
		}
	}
send:
	for {
		select {
		case <-cl.q:
			return
		case dest = <-cl.newDest:
		case dest.C <- Action{dest.CID, m}:
			goto recv
		}
	}
}

func (cl Client) readFromConn() {
	for {
		mt, bs, err := cl.conn.ReadMessage()
		if err != nil {
			fmt.Println("err readFromConn:", err)
			cl.rx <- Close{}
			cl.conn.Close()
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		m, err := UnmarshalJSONToServer(bs)
		if err != nil {
			continue
		}
		cl.rx <- m
	}
}

func (cl Client) writeToConn() {
	ticker := time.NewTicker(PingPeriod)
	var err error
	for {
		select {
		case m, ok := <-cl.tx:
			if !ok {
				ticker.Stop()
				return
			}
			err = cl.conn.WriteJSON(m)
		case <-ticker.C:
			err = cl.conn.WriteMessage(ws.PingMessage, []byte{})
		}
		if err != nil {
			fmt.Println("err writeToConn:", err)
		}
	}
}
