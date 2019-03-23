package main

func runLobbyMap(rx chan ToLobbyMap) {
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	games := make(map[GID]Game)
	next := GID(1)

	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lb := range lobbies {
			ret = append(ret, lb)
		}
		return ret
	}

	broadcastLobbyChoosing := func() {
		msg := LobbyChoices{lobbiesList()}
		for _, cl := range clients.M {
			cl.tx <- msg
		}
	}

	for {
		select {
		case m := <-rx:
			switch m := m.(type) {
			case ClientLeave:
				clients.Add(m.Client)
				m.Client.tx <- LobbyChoices{lobbiesList()}
			case LobbyClose:
				for _, cl := range m.Clients {
					clients.Add(cl)
				}
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			case GameCreate:
				games[m.GID] = NewGame(m.GID, m.Clients, rx)
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			case GameClose:
				delete(games, m.GID)
				msg := LobbyChoices{lobbiesList()}
				for _, cl := range m.Clients {
					clients.Add(cl)
					cl.tx <- msg
				}
			}
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case LobbyChoose:
				lb, ok := lobbies[ts.GID]
				if !ok {
					continue
				}
				lb.tx <- clients.Rm(cid)
			case LobbyCreate:
				gid := next
				next++
				lobbies[gid] = NewLobby(gid, clients.Rm(cid), rx)
				broadcastLobbyChoosing()
			}
		}
	}
}
