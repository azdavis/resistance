package main

import (
	"log"

	ws "github.com/gorilla/websocket"
)

// SendBlocks allows neither sending nor receiving.
var SendBlocks = make(chan<- Action)

// Client is a player of the game. It contains the CID, name, and the way to
// communicate with the actual person represented by this Client.
type Client struct {
	CID                        // unique, never 0
	Name    string             // if "", no name
	tx      chan ToClient      // over the websocket
	rx      chan ToServer      // over the websocket
	acCh    chan<- Action      // everything from rx gets piped here
	newAcCh chan chan<- Action // what to update acCh to
	ackAcCh chan struct{}      // after updating acCh
	q       chan struct{}      // on close
	conn    *ws.Conn           // the websocket
}

// NewClient returns a new client. It starts goroutines to read from and write
// to the given websocket connection. The client return will have CID 0, but
// this should be set to something else immediately.
func NewClient(conn *ws.Conn) *Client {
	cl := &Client{
		CID:     0,
		Name:    "",
		tx:      make(chan ToClient, 3),
		rx:      make(chan ToServer, 3),
		acCh:    SendBlocks,
		newAcCh: make(chan chan<- Action),
		ackAcCh: make(chan struct{}),
		q:       make(chan struct{}),
		conn:    conn,
	}
	go cl.manageAcCh()
	if conn != nil {
		go cl.readFromConn()
		go cl.writeToConn()
	}
	return cl
}

// Close quits the write goroutine. It should be called exactly once, when a
// Close{} ToServer is received from this client. No one else should be reading
// from rx or writing to tx when this is called.
func (cl *Client) Close() {
	close(cl.tx)
	close(cl.q)
}

// Kill kills this Client. It should be called exactly once. No one else should
// be reading from rx or writing to tx when this is called.
func (cl *Client) Kill() {
	if cl.conn != nil {
		cl.conn.Close()
	}
	cl.Close()
}

// SendOn updates acCh.
func (cl *Client) SendOn(acCh chan<- Action) {
	cl.newAcCh <- acCh
	<-cl.ackAcCh
}

func (cl *Client) send(ts ToServer) {
	for {
		select {
		case <-cl.q:
			return
		case acCh := <-cl.newAcCh:
			cl.acCh = acCh
			cl.ackAcCh <- struct{}{}
		case cl.acCh <- Action{cl.CID, ts}:
			return
		}
	}
}

func (cl *Client) manageAcCh() {
	for {
		select {
		case <-cl.q:
			return
		case acCh := <-cl.newAcCh:
			cl.acCh = acCh
			cl.ackAcCh <- struct{}{}
		case ts := <-cl.rx:
			cl.send(ts)
		}
	}
}

// readFromConn reads from the conn, tries to parse the message, and if
// successful, sends the ToServer.
func (cl *Client) readFromConn() {
	log.Println("enter readFromConn", cl.CID)
	for {
		mt, bs, err := cl.conn.ReadMessage()
		if err != nil {
			log.Println("err readFromConn", cl.CID, err)
			cl.rx <- Close{}
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

// writeToConn sends every ToClient from tx over the websocket. See NewClient.
func (cl *Client) writeToConn() {
	log.Println("enter writeToConn", cl.CID)
	defer log.Println("exit writeToConn", cl.CID)
	for m := range cl.tx {
		err := cl.conn.WriteJSON(m)
		if err != nil {
			log.Println("err writeToConn", cl.CID, err)
		}
	}
}
