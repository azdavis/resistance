package main

import (
	"log"
)

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
	leader *Client,
	send chan *Client,
	done chan PID,
) *Party {
	clients := NewClientMap()
	clients.Add(leader)
	p := &Party{
		PID:     pid,
		leader:  leader.CID,
		done:    done,
		send:    send,
		recv:    make(chan *Client),
		clients: clients,
		started: false,
	}
	p.broadcastPartyWaiting()
	go p.run()
	return p
}

// LeaderName returns the name of the leader of this party.
func (p *Party) LeaderName() string {
	return p.clients.M[p.leader].name
}

func (p *Party) clientsInfo() []ClientInfo {
	ret := make([]ClientInfo, 0, len(p.clients.M))
	for cid, cl := range p.clients.M {
		ret = append(ret, ClientInfo{cid, cl.name})
	}
	return ret
}

func (p *Party) broadcastPartyWaiting() {
	clientInfo := p.clientsInfo()
	for cid, cl := range p.clients.M {
		cl.send <- PartyWaiting{
			Self:    cid,
			Leader:  p.leader,
			Clients: clientInfo,
		}
	}
}

func (p *Party) close() {
	log.Println("exit run", p.PID)
	p.done <- p.PID
}

func (p *Party) run() {
	log.Println("enter run", p.PID)
	defer p.close()
	for {
		select {
		case cl := <-p.recv:
			p.clients.Add(cl)
			p.broadcastPartyWaiting()
		case ac := <-p.clients.C:
			cid := ac.CID
			switch ac.ToServer.(type) {
			case Close:
				p.clients.Rm(cid).Close()
				if cid == p.leader || p.started {
					return
				}
				p.broadcastPartyWaiting()
			case PartyLeave:
				if cid == p.leader {
					return
				}
				p.send <- p.clients.Rm(cid)
				p.broadcastPartyWaiting()
			}
		}
	}
}
