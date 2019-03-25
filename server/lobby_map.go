package main

import (
	"sort"
)

func runLobbyMap(rx chan ToLobbyMap, q <-chan struct{}) {
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	games := make(map[GID]Game)
	next := GID(1)

	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lb := range lobbies {
			ret = append(ret, lb)
		}
		// TODO bad perf
		sort.Slice(ret, func(i, j int) bool { return ret[i].GID < ret[j].GID })
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
		case <-q:
			clients.CloseAll()
			return
		case m := <-rx:
			switch m := m.(type) {
			case ClientAdd:
				clients.Add(m.Client)
				m.Client.tx <- LobbyChoices{lobbiesList()}
			case ClientReconnect:
				g, ok := games[m.GID]
				if ok {
					g.tx <- m.Client
				} else {
					clients.Add(m.Client)
					m.Client.tx <- LobbyChoices{lobbiesList()}
				}
			case LobbyClose:
				for _, cl := range m.Clients {
					clients.Add(cl)
				}
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			case GameCreate:
				games[m.GID] = NewGame(m.GID, m.Clients, rx, q)
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			case GameClose:
				delete(games, m.GID)
				m.EndGame.Lobbies = lobbiesList()
				for _, cl := range m.Clients {
					clients.Add(cl)
					cl.tx <- m.EndGame
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
				lobbies[gid] = NewLobby(gid, clients.Rm(cid), rx, q)
				broadcastLobbyChoosing()
			}
		}
	}
}
