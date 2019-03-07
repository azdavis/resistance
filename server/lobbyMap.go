package main

func runLobbyMap(rx chan *Client) {
	const chLen = 5
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	nextGID := GID(1)
	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lobby := range lobbies {
			ret = append(ret, lobby)
		}
		return ret
	}
	toLobbyMap := make(chan LobbyMsg)
	broadcastLobbies := func() {
		msg := LobbyChoosing{Lobbies: lobbiesList()}
		for _, cl := range clients.M {
			cl.tx <- msg
		}
	}
	for {
		select {
		case cl := <-rx:
			clients.Add(cl)
			cl.tx <- LobbyChoosing{Lobbies: lobbiesList()}
		case m := <-toLobbyMap:
			for _, cl := range m.Clients {
				clients.Add(cl)
			}
			if m.Close {
				delete(lobbies, m.GID)
			}
			broadcastLobbies()
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
				broadcastLobbies()
			}
		}
	}
}
