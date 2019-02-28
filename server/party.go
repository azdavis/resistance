package main

// PID is a unique identifier for a Party.
type PID uint64

// Party represents a group of clients all playing together. A Party is totally
// isolated from all other Parties. When the Party is disbanded, all clients
// inside will be sent along send.
type Party struct {
	name   string         // set by leader
	send   chan<- *Client // outgoing clients
	leader *Client        // if leader leaves, party disbands
}

// NewParty returns a new Party. The given Client is the leader of the Party.
func NewParty(name string, send chan<- *Client, leader *Client) *Party {
	p := &Party{
		name:   name,
		send:   send,
		leader: leader,
	}
	return p
}

func (p *Party) run() {}
