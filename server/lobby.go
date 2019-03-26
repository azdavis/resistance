package main

// Lobby represents a group of clients all waiting for the same game to
// start.
type Lobby struct {
	GID                     // unique
	Leader string           // leader name
	tx     chan<- CIDClient // from runLobbyMap to this
}

// NewLobby returns a new Lobby.
func NewLobby(
	gid GID,
	leader CIDClient,
	tx chan<- ToLobbyMap,
	q <-chan struct{},
) Lobby {
	// if this channel is to be buffered, it must be drained when exiting from
	// runLobby, and such draining must only occur after we've sent a message to
	// runLobbyMap that will ensure no further messages get sent on this channel.
	rxLobbyMap := make(chan CIDClient)
	lb := Lobby{
		GID:    gid,
		Leader: "",
		tx:     rxLobbyMap,
	}
	go runLobby(gid, leader, tx, rxLobbyMap, q)
	return lb
}

func runLobby(
	gid GID,
	leader CIDClient,
	tx chan<- ToLobbyMap,
	rx <-chan CIDClient,
	q <-chan struct{},
) {
	// whenever sending on tx, must also select on rx and q to prevent deadlock.

	clients := NewClientMap()
	clients.Add(leader.CID, leader.Client)

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
			clients.CloseAll()
			return
		case cl := <-rx:
			clients.Add(cl.CID, cl.Client)
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
				msg := ClientAdd{cid, clients.Rm(cid)}
			inner:
				for {
					select {
					case <-q:
						clients.CloseAll()
						return
					case cl := <-rx:
						clients.Add(cl.CID, cl.Client)
					case tx <- msg:
						break inner
					}
				}
				broadcastLobbyWaiting()
			case GameStart:
				if cid != leader.CID || !OkGameSize(len(clients.M)) {
					continue
				}
				select {
				case <-q:
					clients.CloseAll()
				case cl := <-rx:
					clients.Add(cl.CID, cl.Client)
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
	clients.DisconnectAll()
	for {
		select {
		case <-q:
			clients.CloseAll()
			return
		case cl := <-rx:
			clients.AddNoSend(cl.CID, cl.Client)
		case tx <- LobbyClose{gid, clients.M}:
			return
		}
	}
}
