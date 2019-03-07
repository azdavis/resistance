package main

import (
	"log"
)

// Lobby represents a group of clients all playing together. A Lobby is totally
// isolated from all other Lobbies.
//
// A Lobby always has at least one client inside. If the leader leaves, the
// lobby disbands. New clients who want to join will arrive on rx. However, the
// lobby will only accept these new clients if the game has not yet started. The
// leader decides when to start the game.
//
// If a single client wants to leave the lobby, send them back to the lobbyMap
// on tx. The tx and rx channels are only to be used before the game has
// started.
//
// A Lobby can disband itself by sending its own GID along done. Once it does
// this, it should stop modifying itself (i.e., it should exit from run), since
// the LobbyMap which contains a pointer to this Lobby will receive along done
// and will start cleaning up the lobby.
type Lobby struct {
	GID                     // unique
	leader  CID             // controls when game starts
	name    string          // leader name
	done    chan<- GID      // send own GID when lobby disbands
	tx      chan<- *Client  // outgoing clients
	rx      chan *Client    // incoming clients
	clients *ClientMap      // clients in the lobby (includes leader)
	started bool            // whether the game has started
	start   chan<- struct{} // signal when the game has started
}

// NewLobby returns a new Lobby.
func NewLobby(
	gid GID,
	leader *Client,
	tx chan<- *Client,
	done chan<- GID,
	start chan<- struct{},
) *Lobby {
	clients := NewClientMap()
	clients.Add(leader)
	p := &Lobby{
		GID:     gid,
		leader:  leader.CID,
		name:    leader.name,
		done:    done,
		tx:      tx,
		rx:      make(chan *Client),
		clients: clients,
		started: false,
		start:   start,
	}
	p.broadcastLobbyWaiting()
	go p.run()
	return p
}

// LeaderName returns the name of the leader of this lobby.
func (p *Lobby) LeaderName() string {
	return p.name
}

// clientsInfo returns info about the Clients in this Lobby.
func (p *Lobby) clientsInfo() []ClientInfo {
	ret := make([]ClientInfo, 0, len(p.clients.M))
	for cid, cl := range p.clients.M {
		ret = append(ret, ClientInfo{cid, cl.name})
	}
	return ret
}

// broadcastLobbyWaiting broadcasts a LobbyWaiting message to every Client in
// this Lobby.
func (p *Lobby) broadcastLobbyWaiting() {
	clientInfo := p.clientsInfo()
	for cid, cl := range p.clients.M {
		cl.tx <- LobbyWaiting{
			Self:    cid,
			Leader:  p.leader,
			Clients: clientInfo,
		}
	}
}

// close signals the LobbyMap to clean up this Lobby. It should only be called
// from run.
func (p *Lobby) close() {
	log.Println("exit run", p.GID)
	p.done <- p.GID
}

// run runs the Lobby. When run returns, any remaining Clients are absorbed into
// the LobbyMap.
func (p *Lobby) run() {
	log.Println("enter run", p.GID)
	defer p.close()
	for {
		select {
		case cl := <-p.rx:
			if p.started {
				continue
			}
			p.clients.Add(cl)
			p.broadcastLobbyWaiting()
		case ac := <-p.clients.C:
			cid := ac.CID
			switch ac.ToServer.(type) {
			case Close:
				p.clients.Rm(cid).Close()
				if cid == p.leader || p.started {
					return
				}
				p.broadcastLobbyWaiting()
			case LobbyLeave:
				if cid == p.leader {
					return
				}
				p.tx <- p.clients.Rm(cid)
				p.broadcastLobbyWaiting()
			case GameStart:
				if cid != p.leader || p.started {
					continue
				}
				p.started = true
				p.start <- struct{}{}
			}
		}
	}
}
