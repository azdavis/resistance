package main

import (
	"sort"
)

func runLobbyMap(rx chan ToLobbyMap, q <-chan struct{}) {
	clients := NewClientMap()
	lobbies := make(map[GID]Lobby)
	games := make(map[GID]Game)
	names := make(map[CID]string)
	nextGID := GID(1)

	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lb := range lobbies {
			ret = append(ret, lb)
		}
		// TODO bad perf
		sort.Slice(ret, func(i, j int) bool { return ret[i].GID < ret[j].GID })
		return ret
	}

	mkClientAdd := func(cid CID) ClientAdd {
		ca := ClientAdd{cid, clients.Rm(cid), names[cid]}
		delete(names, cid)
		return ca
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
				clients.Add(m.CID, m.Client)
				names[m.CID] = m.Name
				m.Client.tx <- LobbyChoices{lobbiesList()}
			case ClientReconnect:
				g, ok := games[m.GID]
				if ok {
					g.tx <- CIDClient{m.CID, m.Client}
				} else {
					clients.Add(m.CID, m.Client)
					m.Client.tx <- LobbyChoices{lobbiesList()}
				}
			case LobbyClose:
				if m.MakeGame {
					games[m.GID] = NewGame(m.GID, m.Clients, m.Names, rx, q)
				} else {
					for cid, cl := range m.Clients.M {
						clients.Add(cid, cl)
					}
					for cid, name := range m.Names {
						names[cid] = name
					}
				}
				delete(lobbies, m.GID)
				broadcastLobbyChoosing()
			case GameClose:
				m.EndGame.Lobbies = lobbiesList()
				for cid, cl := range m.Clients.M {
					clients.Add(cid, cl)
					cl.tx <- m.EndGame
				}
				for cid, name := range m.Names {
					names[cid] = name
				}
				delete(games, m.GID)
			}
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				delete(names, cid)
			case LobbyChoose:
				lb, ok := lobbies[ts.GID]
				if !ok {
					continue
				}
				lb.tx <- mkClientAdd(cid)
			case LobbyCreate:
				lobbies[nextGID] = NewLobby(nextGID, mkClientAdd(cid), rx, q)
				nextGID++
				broadcastLobbyChoosing()
			}
		}
	}
}
