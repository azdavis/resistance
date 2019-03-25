package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

// Client is a player of the game. It contains the CID, name, and the way to
// communicate with the actual person represented by this Client.
type Client struct {
	CID                // unique, never 0
	Name string        // if "", no name
	tx   chan ToClient // over the websocket
	rx   chan ToServer // over the websocket
	conn *ws.Conn      // the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The client return will have CID 0, but
// this should be set to something else immediately.
func NewClient(conn *ws.Conn) *Client {
	cl := &Client{
		CID:  0,
		Name: "",
		tx:   make(chan ToClient),
		rx:   make(chan ToServer),
		conn: conn,
	}
	go cl.doRx()
	go cl.doTx()
	return cl
}

// Close quits the write goroutine. It should be called exactly once, when a
// Close{} ToServer is received from this client. No one else should be reading
// from rx or writing to tx when this is called.
func (cl *Client) Close() {
	close(cl.tx)
}

// Kill kills this Client. It should be called exactly once. No one else should
// be reading from rx or writing to tx when this is called.
func (cl *Client) Kill() {
	if cl.conn == nil {
		close(cl.rx)
	} else {
		cl.conn.Close()
	}
	for range cl.rx {
	}
	cl.Close()
}

// doRx reads from the conn, tries to parse the message, and if successful,
// sends the ToServer over rx.
func (cl *Client) doRx() {
	log.Println("enter doRx", cl.CID)
	defer log.Println("exit doRx", cl.CID)
	for {
		mt, bs, err := cl.conn.ReadMessage()
		if err != nil {
			log.Println("err doRx", cl.CID, err)
			cl.rx <- Close{}
			close(cl.rx)
			cl.conn.Close()
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

// doTx sends every ToClient from tx over the websocket. See NewClient.
func (cl *Client) doTx() {
	log.Println("enter doTx", cl.CID)
	defer log.Println("exit doTx", cl.CID)
	for m := range cl.tx {
		err := cl.conn.WriteJSON(m)
		if err != nil {
			log.Println("err doTx", cl.CID, err)
		}
	}
}
