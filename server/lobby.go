package main

import (
	"log"
)

// Lobby represents a group of clients all waiting for the same game to
// start.
type Lobby struct {
	GID                   // unique
	Leader string         // leader name
	tx     chan<- *Client // from runLobbyMap to this
}

// NewLobby returns a new Lobby.
func NewLobby(
	gid GID,
	leader *Client,
	tx chan<- ToLobbyMap,
	q <-chan struct{},
) Lobby {
	// if this channel is to be buffered, it must be drained when exiting from
	// runLobby, and such draining must only occur after we've sent a message to
	// runLobbyMap that will ensure no further messages get sent on this channel.
	rxLobbyMap := make(chan *Client)
	lb := Lobby{
		GID:    gid,
		Leader: leader.Name,
		tx:     rxLobbyMap,
	}
	go runLobby(gid, leader, tx, rxLobbyMap, q)
	return lb
}

func runLobby(
	gid GID,
	leader *Client,
	tx chan<- ToLobbyMap,
	rx <-chan *Client,
	q <-chan struct{},
) {
	// whenever sending on tx, must also select on rx and q to prevent deadlock.
	log.Println("enter runLobby", gid)
	defer log.Println("exit runLobby", gid)

	clients := NewClientMap()
	clients.Add(leader)

	broadcastLobbyWaiting := func() {
		cs := clients.ToList()
		for _, cl := range clients.M {
			cl.tx <- CurrentLobby{gid, leader.CID, cs}
		}
	}
	broadcastLobbyWaiting()

	for {
		select {
		case <-q:
			clients.KillAll()
			return
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
				msg := ClientAdd{clients.Rm(cid)}
				select {
				case <-q:
					clients.KillAll()
					return
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
				case <-q:
					clients.KillAll()
				case cl := <-rx:
					clients.Add(cl)
					// allow leader to re-verify whether the game should be started.
					broadcastLobbyWaiting()
					continue
				case tx <- GameCreate{gid, clients}:
				}
				return
			}
		}
	}

out:
	cs := clients.Clear()
	for {
		select {
		case <-q:
			clients.KillAll()
			return
		case cl := <-rx:
			cs = append(cs, cl)
		case tx <- LobbyClose{gid, cs}:
			return
		}
	}
}
