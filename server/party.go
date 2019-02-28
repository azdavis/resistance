package main

// PID is a unique identifier for a Party.
type PID uint64

// Party represents a group of clients all playing together. A Party is totally
// isolated from all other Parties.
type Party struct {
	leader *Client // the leader
}

// NewParty returns a new Party. The given Client is the leader of the Party.
func NewParty(leader *Client) *Party {
	p := &Party{
		leader: leader,
	}
	return p
}

func (p *Party) run() {}
