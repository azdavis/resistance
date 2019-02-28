package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

// CID is a unique identifier for a Client.
type CID uint64

// Client is a player of the game. It contains the CID, game information, and
// the way to communicate with the actual person represented by this Client.
// Close should be called after a Close{} is received on recv.
type Client struct {
	id    CID         // unique, never 0
	pid   PID         // if 0, no party
	name  string      // if "", no name
	isSpy bool        // if false, is resistance
	send  chan State  // send the current State over the websocket
	recv  chan Action // receive an Action over the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The id should not be in use by any other
// client. send should be closed when this Client will no longer be used.
func NewClient(conn *ws.Conn, id CID) *Client {
	const chLen = 3
	cl := &Client{
		id:    id,
		pid:   0,
		name:  "",
		isSpy: false,
		send:  make(chan State, chLen),
		recv:  make(chan Action, chLen),
	}
	go cl.recvFrom(conn)
	go cl.sendTo(conn)
	return cl
}

// Close quits the write goroutine. It should be called only once.
func (cl *Client) Close() {
	close(cl.send)
}

// recvFrom reads from the conn, tries to parse the message, and if successful,
// sends the Action over recv.
func (cl *Client) recvFrom(conn *ws.Conn) {
	for {
		mt, bs, err := conn.ReadMessage()
		if err != nil {
			log.Println("recvFrom", cl.id, err)
			// no further actions will be sent on recv. however, do not close recv,
			// since we may send garbage actions to listeners. only send the Close
			// Action.
			cl.recv <- Close{}
			conn.Close()
			log.Println("exit recvFrom", cl.id)
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

// sendTo sends every message from send over the websocket. See NewClient.
func (cl *Client) sendTo(conn *ws.Conn) {
	for m := range cl.send {
		err := conn.WriteJSON(m)
		if err != nil {
			log.Println("sendTo", cl.id, err)
		}
	}
	log.Println("exit sendTo", cl.id)
}
