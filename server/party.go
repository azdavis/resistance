package main

// PID is a unique identifier for a Party.
type PID uint64

// Party represents a group of clients all playing together. A Party is totally
// isolated from all other Parties.
//
// A Party always has at least one client inside. If the leader leaves, the
// party disbands. New clients who want to join will arrive on recv. However,
// the party will only accept these new clients if the game has not yet started.
// The leader decides when to start the game.
//
// A Party can disband itself by sending its own PID along done. Once it does
// this, it should stop modifying itself (i.e., it should exit from run), since
// the Lobby which contains a pointer to this Party will receive along done and
// will start cleaning up the party.
type Party struct {
	id      PID          // unique
	name    string       // set by leader
	leader  CID          // controls when game starts, can unilaterally disband
	done    chan PID     // send own id when party disbands
	recv    chan *Client // incoming clients
	clients *ClientMap   // clients in the party (includes leader)
}

// NewParty returns a new Party.
func NewParty(
	id PID,
	name string,
	leader *Client,
	done chan PID,
) *Party {
	clients := NewClientMap()
	clients.Add(leader)
	p := &Party{
		id:      id,
		name:    name,
		leader:  leader.id,
		done:    done,
		recv:    make(chan *Client),
		clients: clients,
	}
	go p.run()
	return p
}

func (p *Party) run() {
	for {
		select {
		case cl := <-p.recv:
			p.clients.Add(cl)
		case cidAc := <-p.clients.C:
			id := cidAc.CID
			switch cidAc.Action.(type) {
			case Close:
				p.clients.Rm(id).Close()
				if id == p.leader {
					p.done <- p.id
					return
				}
			}
		}
	}
}
