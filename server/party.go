package main

// PID is a unique identifier for a Party.
type PID uint64

// Party represents a group of clients all playing together. A Party is totally
// isolated from all other Parties. When the Party is disbanded, all clients
// inside will be sent along send.
type Party struct {
	send   chan<- *Client // outgoing clients
	leader *Client        // the leader
}

// NewParty returns a new Party. The given Client is the leader of the Party.
func NewParty(send chan<- *Client, leader *Client) *Party {
	p := &Party{
		send:   send,
		leader: leader,
	}
	return p
}

func (p *Party) run() {}
