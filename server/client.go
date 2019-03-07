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
// Close should be called after a Close{} is received on rx.
type Client struct {
	CID                // unique, never 0
	name string        // if "", no name
	tx   chan ToClient // over the websocket
	rx   chan ToServer // over the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The CID should not be in use by any other
// client. tx should be closed when this Client will no longer be used.
func NewClient(conn *ws.Conn, cid CID) *Client {
	cl := &Client{
		CID:  cid,
		name: "",
		tx:   make(chan ToClient, clientChLen),
		rx:   make(chan ToServer, clientChLen),
	}
	go cl.recvFrom(conn)
	go cl.sendTo(conn)
	return cl
}

// Close quits the write goroutine. It should be called exactly once, when a
// Close{} ToServer is received from this client.
func (cl *Client) Close() {
	close(cl.tx)
}

// recvFrom reads from the conn, tries to parse the message, and if successful,
// sends the ToServer over rx.
func (cl *Client) recvFrom(conn *ws.Conn) {
	log.Println("enter recvFrom", cl.CID)
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("err recvFrom", cl.CID, err)
			// no further ToServers will be sent on rx. however, do not close rx,
			// since we may tx garbage ToServers to listeners.
			cl.rx <- Close{}
			conn.Close()
			log.Println("exit recvFrom", cl.CID)
			return
		}
		if mt != ws.TextMessage {
			continue
		}
		ts, err := UnmarshalJSONToServer(bs)
		if err != nil {
			continue
		}
		cl.rx <- ts
	}
}

// sendTo sends every ToClient from tx over the websocket. See NewClient.
func (cl *Client) sendTo(conn *ws.Conn) {
	log.Println("enter sendTo", cl.CID)
	for m := range cl.tx {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("err sendTo", cl.CID, err)
		}
	}
	log.Println("exit sendTo", cl.CID)
}
