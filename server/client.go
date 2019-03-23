package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

// Client is a player of the game. It contains the CID, game information, and
// the way to communicate with the actual person represented by this Client.
// Close should be called after a Close{} is received on rx.
type Client struct {
	CID                // unique, never 0
	Name string        // if "", no name
	tx   chan ToClient // over the websocket
	rx   chan ToServer // over the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The CID should not be in use by any other
// client. tx should be closed when this Client will no longer be used.
func NewClient(conn *ws.Conn, cid CID) *Client {
	const chLen = 3
	cl := &Client{
		CID:  cid,
		Name: "",
		tx:   make(chan ToClient, chLen),
		rx:   make(chan ToServer, chLen),
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
	defer log.Println("exit recvFrom", cl.CID)
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("err recvFrom", cl.CID, err)
			cl.rx <- Close{}
			close(cl.rx)
			conn.Close()
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
	defer log.Println("exit sendTo", cl.CID)
	for m := range cl.tx {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("err sendTo", cl.CID, err)
		}
	}
}
