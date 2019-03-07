package main

import (
	"log"
)

// Lobby represents a group of clients all waiting for the same game to
// start.
type Lobby struct {
	GID                   // unique
	Leader string         // leader name
	tx     chan<- *Client // incoming clients
}

// NewLobby returns a new Lobby.
func NewLobby(
	gid GID,
	leader *Client,
	toLobbyMap chan<- LobbyMsg,
) Lobby {
	clients := NewClientMap()
	clients.Add(leader)
	toLobby := make(chan *Client)
	lb := Lobby{
		GID:    gid,
		Leader: leader.name,
		tx:     toLobby,
	}
	go lb.run(gid, leader, toLobbyMap, toLobby)
	return lb
}

// run runs the Lobby. When run returns, any remaining Clients are absorbed into
// the LobbyMap.
func (lb Lobby) run(
	gid GID,
	leader *Client,
	tx chan<- LobbyMsg,
	rx <-chan *Client,
) {
	clients := NewClientMap()
	clients.Add(leader)
	clientsInfo := func() []ClientInfo {
		ret := make([]ClientInfo, 0, len(clients.M))
		for cid, cl := range clients.M {
			ret = append(ret, ClientInfo{cid, cl.name})
		}
		return ret
	}
	broadcastLobbyWaiting := func() {
		clientInfo := clientsInfo()
		for cid, cl := range clients.M {
			cl.tx <- LobbyWaiting{
				Self:    cid,
				Leader:  leader.CID,
				Clients: clientInfo,
			}
		}
	}
	broadcastLobbyWaiting()
	log.Println("enter run", lb.GID)
	for {
		select {
		case cl := <-rx:
			clients.Add(cl)
			broadcastLobbyWaiting()
		case ac := <-clients.C:
			cid := ac.CID
			switch ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				if cid == leader.CID {
					return
				}
				broadcastLobbyWaiting()
			case LobbyLeave:
				if cid == leader.CID {
					return
				}
				lb.tx <- clients.Rm(cid)
				broadcastLobbyWaiting()
			case GameStart:
				if cid != leader.CID {
					continue
				}
			}
		}
	}
}
