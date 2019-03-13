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
func NewLobby(gid GID, leader *Client, toLobbyMap chan<- LobbyMsg) Lobby {
	// if this channel is to be buffered, it must be drained when exiting from
	// runLobby.
	toLobby := make(chan *Client)
	lb := Lobby{
		GID:    gid,
		Leader: leader.Name,
		tx:     toLobby,
	}
	go runLobby(gid, leader, toLobbyMap, toLobby)
	return lb
}

func runLobby(gid GID, leader *Client, tx chan<- LobbyMsg, rx <-chan *Client) {
	// keep in sync with LobbyWaiter.tsx and game.go
	// whenever sending on tx, must also select with rx to prevent deadlock.
	log.Println("enter runLobby", gid)
	defer log.Println("exit runLobby", gid)

	clients := NewClientMap()
	clients.Add(leader)

	broadcastLobbyWaiting := func() {
		cs := clients.ToList()
		for cid, cl := range clients.M {
			cl.tx <- CurrentLobby{cid, leader.CID, cs}
		}
	}
	broadcastLobbyWaiting()

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
				if cid != leader.CID || !OkGameSize(len(clients.M)) {
					continue
				}
				select {
				case cl := <-rx:
					clients.Add(cl)
					// allow leader to re-verify whether the game should be started.
					broadcastLobbyWaiting()
					continue
				case tx <- LobbyMsg{gid, true, []*Client{}}:
				}
				go runGame(gid, tx, clients)
				return
			}
		}
	}

out:
	cs := clients.Clear()
	select {
	case cl := <-rx:
		cs = append(cs, cl)
	case tx <- LobbyMsg{gid, true, cs}:
	}
}
