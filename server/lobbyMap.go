package main

func runLobbyMap(rx chan *Client) {
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	toLobbyMap := make(chan LobbyMsg)
	nextGID := GID(1)

	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lobby := range lobbies {
			ret = append(ret, lobby)
		}
		return ret
	}

	broadcastLobbyChoosing := func() {
		msg := LobbyChoices{Lobbies: lobbiesList()}
		for _, cl := range clients.M {
			cl.tx <- msg
		}
	}

	for {
		select {
		case cl := <-rx:
			clients.Add(cl)
			cl.tx <- LobbyChoices{Lobbies: lobbiesList()}
		case m := <-toLobbyMap:
			for _, cl := range m.Clients {
				clients.Add(cl)
			}
			if m.Close {
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			} else {
				msg := LobbyChoices{Lobbies: lobbiesList()}
				for _, cl := range m.Clients {
					cl.tx <- msg
				}
			}
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case LobbyChoose:
				lobby, ok := lobbies[ts.GID]
				if !ok {
					continue
				}
				lobby.tx <- clients.Rm(cid)
			case LobbyCreate:
				gid := nextGID
				nextGID++
				lobbies[gid] = NewLobby(gid, clients.Rm(cid), toLobbyMap)
				broadcastLobbyChoosing()
			}
		}
	}
}
