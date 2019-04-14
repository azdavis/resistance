package main

// Lobby represents a group of clients all waiting for the same game to
// start.
type Lobby struct {
	GID                       // unique
	Leader string             // leader name
	tx     chan<- NamedClient // from runLobbyMap to this
}

// NewLobby returns a new Lobby.
func NewLobby(
	gid GID,
	leader NamedClient,
	tx chan<- SrvMsg,
	q <-chan struct{},
) Lobby {
	// if this channel is to be buffered, it must be drained when exiting from
	// runLobby, and such draining must only occur after we've sent a message to
	// runLobbyMap that will ensure no further messages get sent on this channel.
	rxLobbyMap := make(chan NamedClient)
	go runLobby(gid, leader, tx, rxLobbyMap, q)
	return Lobby{
		GID:    gid,
		Leader: leader.Name,
		tx:     rxLobbyMap,
	}
}

func runLobby(
	gid GID,
	leader NamedClient,
	tx chan<- SrvMsg,
	rx <-chan NamedClient,
	q <-chan struct{},
) {
	// whenever sending on tx, must also select on rx and q to prevent deadlock.

	clients := NewClientMap()
	clients.Add(leader.CID, leader.Client)

	names := make(map[CID]string)
	names[leader.CID] = leader.Name

	broadcastLobbyWaiting := func() {
		cs := clients.ToList(names)
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
			names[cl.CID] = cl.Name
			broadcastLobbyWaiting()
		case m := <-clients.C:
			cid := m.CID
			switch m.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				delete(names, cid)
				if cid == leader.CID {
					goto out
				}
				broadcastLobbyWaiting()
			case LobbyLeave:
				if cid == leader.CID {
					goto out
				}
				msg := NamedClient{cid, clients.Rm(cid), names[cid]}
				delete(names, cid)
			inner:
				for {
					select {
					case <-q:
						clients.CloseAll()
						return
					case cl := <-rx:
						clients.Add(cl.CID, cl.Client)
						names[cl.CID] = cl.Name
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
					names[cl.CID] = cl.Name
					// allow leader to re-verify whether the game should be started.
					broadcastLobbyWaiting()
					continue
				case tx <- LobbyClose{true, gid, clients, names}:
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
			names[cl.CID] = cl.Name
		case tx <- LobbyClose{false, gid, clients, names}:
			return
		}
	}
}
