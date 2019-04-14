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
	tx   chan ToClient // orders for the client
	rx   chan ToServer // requests from the client
	d    chan Dest     // what to update the ultimate destination of rx to
	q    chan struct{} // close on Close
	conn *ws.Conn      // the websocket
}

// NewClient returns a new Client.
func NewClient(conn *ws.Conn) Client {
	tx := make(chan ToClient, 3)
	rx := make(chan ToServer, 3)
	d := make(chan Dest)
	q := make(chan struct{})
	cl := Client{tx, rx, d, q, conn}
	go manageDest(q, d, rx)
	if conn != nil {
		go readFromConn(conn, rx)
		go writeToConn(conn, tx)
	}
	return cl
}

// Close cleans up resources for this Client. It should be called exactly once.
// Usually this is called after receiving a Close{} on rx.
func (cl Client) Close() {
	if cl.conn != nil {
		cl.conn.Close()
	}
	close(cl.q)
	close(cl.tx)
}

// RecvTo updates dest. It returns only once dest has been updated.
func (cl Client) RecvTo(dest Dest) {
	cl.d <- dest
}

func manageDest(q <-chan struct{}, d <-chan Dest, rx <-chan ToServer) {
	dest := NullDest
	var m ToServer
recv:
	for {
		select {
		case <-q:
			return
		case dest = <-d:
		case m = <-rx:
			goto send
		}
	}
send:
	for {
		select {
		case <-q:
			return
		case dest = <-d:
		case dest.C <- Action{dest.CID, m}:
			goto recv
		}
	}
}

func readFromConn(conn *ws.Conn, rx chan<- ToServer) {
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("err readFromConn:", err)
			rx <- Close{}
			conn.Close()
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		m, err := UnmarshalJSONToServer(bs)
		if err != nil {
			continue
		}
		rx <- m
	}
}

func writeToConn(conn *ws.Conn, tx <-chan ToClient) {
	ticker := time.NewTicker(PingPeriod)
	var err error
	for {
		select {
		case m, ok := <-tx:
			if !ok {
				ticker.Stop()
				return
			}
			err = conn.WriteJSON(m)
		case <-ticker.C:
			err = conn.WriteMessage(ws.PingMessage, []byte{})
		}
		if err != nil {
			fmt.Println("err writeToConn:", err)
		}
	}
}
