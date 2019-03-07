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
	toLobby := make(chan *Client)
	lb := Lobby{
		GID:    gid,
		Leader: leader.Name,
		tx:     toLobby,
	}
	go runLobby(gid, leader, toLobbyMap, toLobby)
	return lb
}

func runLobby(
	gid GID,
	leader *Client,
	tx chan<- LobbyMsg,
	rx <-chan *Client,
) {
	// whenever sending on tx, must also select with rx to prevent deadlock.
	clients := NewClientMap()
	clients.Add(leader)
	clientsList := func() []*Client {
		ret := make([]*Client, 0, len(clients.M))
		for _, cl := range clients.M {
			ret = append(ret, cl)
		}
		return ret
	}
	clientsInfo := func() []ClientInfo {
		ret := make([]ClientInfo, 0, len(clients.M))
		for cid, cl := range clients.M {
			ret = append(ret, ClientInfo{cid, cl.Name})
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
	log.Println("enter run", gid)
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
					goto out
				}
				broadcastLobbyWaiting()
			case LobbyLeave:
				if cid == leader.CID {
					goto out
				}
				msg := LobbyMsg{gid, false, []*Client{clients.Rm(cid)}}
				select {
				case cl := <-rx:
					clients.Add(cl)
				case tx <- msg:
				}
				broadcastLobbyWaiting()
			case GameStart:
				if cid != leader.CID {
					continue
				}
				// TODO
			}
		}
	}
out:
	select {
	case cl := <-rx:
		clients.Add(cl)
	case tx <- LobbyMsg{gid, true, clientsList()}:
	}
	log.Println("exit run", gid)
}
