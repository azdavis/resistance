package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

const clientChLen = 3

// CID is a unique identifier for a Client.
type CID uint64

// Client is a player of the game. It contains the CID, game information, and
// the way to communicate with the actual person represented by this Client.
// Close should be called after a Close{} is received on recv.
type Client struct {
	CID                 // unique, never 0
	name  string        // if "", no name
	isSpy bool          // if false, is resistance
	send  chan ToClient // over the websocket
	recv  chan ToServer // over the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The CID should not be in use by any other
// client. send should be closed when this Client will no longer be used.
func NewClient(conn *ws.Conn, cid CID) *Client {
	cl := &Client{
		CID:   cid,
		name:  "",
		isSpy: false,
		send:  make(chan ToClient, clientChLen),
		recv:  make(chan ToServer, clientChLen),
	}
	go cl.recvFrom(conn)
	go cl.sendTo(conn)
	return cl
}

// Close quits the write goroutine. It should be called exactly once, when a
// Close{} ToServer is received from this client.
func (cl *Client) Close() {
	close(cl.send)
}

// recvFrom reads from the conn, tries to parse the message, and if successful,
// sends the ToServer over recv.
func (cl *Client) recvFrom(conn *ws.Conn) {
	log.Println("enter recvFrom", cl.CID)
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("err recvFrom", cl.CID, err)
			// no further ToServers will be sent on recv. however, do not close recv,
			// since we may send garbage ToServers to listeners.
			cl.recv <- Close{}
			conn.Close()
			log.Println("exit recvFrom", cl.CID)
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

// sendTo sends every ToClient from send over the websocket. See NewClient.
func (cl *Client) sendTo(conn *ws.Conn) {
	log.Println("enter sendTo", cl.CID)
	for m := range cl.send {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("err sendTo", cl.CID, err)
		}
	}
	log.Println("exit sendTo", cl.CID)
}
