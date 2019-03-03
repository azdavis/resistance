package main

// TODO this provides no way for a client to leave the party before the game
// starts.

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
// If a single client wants to leave the party, send them back to the lobby on
// send. The send and recv channels are only to be used before the game has
// started.
//
// A Party can disband itself by sending its own PID along done. Once it does
// this, it should stop modifying itself (i.e., it should exit from run), since
// the Lobby which contains a pointer to this Party will receive along done and
// will start cleaning up the party.
type Party struct {
	PID                  // unique
	name    string       // set by leader
	leader  CID          // controls when game starts, can unilaterally disband
	done    chan PID     // send own PID when party disbands
	send    chan *Client // outgoing clients
	recv    chan *Client // incoming clients
	clients *ClientMap   // clients in the party (includes leader)
	started bool         // whether the game has started
}

// NewParty returns a new Party.
func NewParty(
	pid PID,
	name string,
	leader *Client,
	send chan *Client,
	done chan PID,
) *Party {
	clients := NewClientMap()
	clients.Add(leader)
	p := &Party{
		PID:     pid,
		name:    name,
		leader:  leader.CID,
		done:    done,
		send:    send,
		recv:    make(chan *Client),
		clients: clients,
		started: false,
	}
	go p.run()
	return p
}

func (p *Party) run() {
	for {
		select {
		case cl := <-p.recv:
			p.clients.Add(cl)
		case ac := <-p.clients.C:
			cid := ac.CID
			switch ac.ToServer.(type) {
			case Close:
				p.clients.Rm(cid).Close()
				if cid == p.leader || p.started {
					p.done <- p.PID
					return
				}
			case PartyLeave:
				if cid == p.leader {
					p.done <- p.PID
				} else {
					p.send <- p.clients.Rm(cid)
				}
			}
		}
	}
}
