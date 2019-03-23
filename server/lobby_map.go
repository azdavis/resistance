package main

func runLobbyMap(rx chan *Client) {
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	toLobbyMap := make(chan LobbyMsg)
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
		case cl := <-rx:
			clients.Add(cl)
			cl.tx <- LobbyChoices{lobbiesList()}
		case m := <-toLobbyMap:
			for _, cl := range m.Clients {
				clients.Add(cl)
			}
			if m.Close {
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			} else {
				msg := LobbyChoices{lobbiesList()}
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
				lb, ok := lobbies[ts.GID]
				if !ok {
					continue
				}
				lb.tx <- clients.Rm(cid)
			case LobbyCreate:
				gid := next
				next++
				lobbies[gid] = NewLobby(gid, clients.Rm(cid), toLobbyMap)
				broadcastLobbyChoosing()
			}
		}
	}
}