package main

import (
	"log"
)

func runLobbyMap(rx chan *Client) {
	const chLen = 5
	log.Println("enter LobbyMap run")
	clients := NewClientMap()
	lobbies := make(map[GID]*Lobby)
	nextGID := GID(1)
	lobbiesInfo := func() []LobbyInfo {
		ret := make([]LobbyInfo, 0, len(lobbies))
		for gid, lobby := range lobbies {
			// TODO race with lobby.run
			if lobby.started {
				continue
			}
			ret = append(ret, LobbyInfo{gid, lobby.LeaderName()})
		}
		return ret
	}
	done := make(chan GID, chLen)
	start := make(chan struct{}, chLen)
	broadcastLobbies := func() {
		msg := LobbyChoosing{Lobbies: lobbiesInfo()}
		for _, cl := range clients.M {
			if cl.name != "" {
				cl.tx <- msg
			}
		}
	}
	for {
		select {
		case cl := <-rx:
			clients.Add(cl)
			if cl.name != "" {
				cl.tx <- LobbyChoosing{Lobbies: lobbiesInfo()}
			}
		case <-start:
			broadcastLobbies()
		case gid := <-done:
			rmLobbyClients := lobbies[gid].clients
			for cid := range rmLobbyClients.M {
				clients.Add(rmLobbyClients.Rm(cid))
			}
			delete(lobbies, gid)
			broadcastLobbies()
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case NameChoose:
				log.Println("NameChoose", cid, ts.Name)
				cl, ok := clients.M[cid]
				if !ok {
					continue
				}
				if !validName(ts.Name) {
					cl.tx <- NameChoosing{Valid: false}
					continue
				}
				cl.name = ts.Name
				cl.tx <- LobbyChoosing{Lobbies: lobbiesInfo()}
			case LobbyChoose:
				log.Println("LobbyChoose", cid, ts.GID)
				lobby, ok := lobbies[ts.GID]
				if !ok {
					continue
				}
				lobby.rx <- clients.Rm(cid)
			case LobbyCreate:
				log.Println("LobbyCreate", cid)
				_, ok := clients.M[cid]
				if !ok {
					continue
				}
				gid := nextGID
				nextGID++
				lobbies[gid] = NewLobby(
					gid,
					clients.Rm(cid),
					rx,
					done,
					start,
				)
				broadcastLobbies()
			}
		}
	}
}
