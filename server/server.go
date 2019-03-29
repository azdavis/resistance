package main

import (
	"sort"
)

// Server is the server.
type Server struct {
	C chan<- SrvMsg
	q chan<- struct{}
}

// NewServer returns a new Server.
func NewServer() *Server {
	C := make(chan SrvMsg, 3)
	q := make(chan struct{})
	go runServer(C, q)
	s := &Server{C, q}
	return s
}

// Close shuts down the Server. It should only be called once.
func (s *Server) Close() {
	close(s.q)
}

func runServer(rx chan SrvMsg, q <-chan struct{}) {
	named := NewClientMap()
	unnamed := NewClientMap()
	lobbies := make(map[GID]Lobby)
	games := make(map[GID]Game)
	names := make(map[CID]string)
	nextGID := GID(1)
	nextCID := CID(1)

	lobbiesList := func() []Lobby {
		ret := make([]Lobby, 0, len(lobbies))
		for _, lb := range lobbies {
			ret = append(ret, lb)
		}
		// TODO bad perf
		sort.Slice(ret, func(i, j int) bool { return ret[i].GID < ret[j].GID })
		return ret
	}

	mkNamedClient := func(cid CID) NamedClient {
		ca := NamedClient{cid, named.Rm(cid), names[cid]}
		delete(names, cid)
		return ca
	}

	broadcastLobbies := func() {
		msg := LobbyChoices{lobbiesList()}
		for _, cl := range named.M {
			cl.tx <- msg
		}
	}

	for {
		select {
		case <-q:
			named.CloseAll()
			unnamed.CloseAll()
			return
		case m := <-rx:
			switch m := m.(type) {
			case Client:
				unnamed.Add(nextCID, m)
				nextCID++
			case NamedClient:
				named.Add(m.CID, m.Client)
				names[m.CID] = m.Name
				m.Client.tx <- LobbyChoices{lobbiesList()}
			case LobbyClose:
				if m.MakeGame {
					games[m.GID] = NewGame(m.GID, m.Clients, m.Names, rx, q)
				} else {
					for cid, cl := range m.Clients.M {
						named.Add(cid, cl)
					}
					for cid, name := range m.Names {
						names[cid] = name
					}
				}
				delete(lobbies, m.GID)
				broadcastLobbies()
			case GameClose:
				m.EndGame.Lobbies = lobbiesList()
				for cid, cl := range m.Clients.M {
					named.Add(cid, cl)
					cl.tx <- m.EndGame
				}
				for cid, name := range m.Names {
					names[cid] = name
				}
				delete(games, m.GID)
			}
		case m := <-named.C:
			cid := m.CID
			switch m := m.ToServer.(type) {
			case Close:
				named.Rm(cid).Close()
				delete(names, cid)
			case LobbyChoose:
				lb, ok := lobbies[m.GID]
				if !ok {
					continue
				}
				lb.tx <- mkNamedClient(cid)
			case LobbyCreate:
				lobbies[nextGID] = NewLobby(nextGID, mkNamedClient(cid), rx, q)
				nextGID++
				broadcastLobbies()
			}
		case m := <-unnamed.C:
			cid := m.CID
			switch m := m.ToServer.(type) {
			case Close:
				unnamed.Rm(cid).Close()
			case Connect:
				unnamed.M[cid].tx <- SetMe{cid}
			case Reconnect:
				g, ok := games[m.GID]
				if ok {
					g.tx <- CIDClient{m.Me, unnamed.Rm(cid)}
				} else {
					unnamed.M[cid].tx <- SetMe{cid}
				}
			case NameChoose:
				if ValidName(m.Name) {
					cl := unnamed.Rm(cid)
					named.Add(cid, cl)
					names[cid] = m.Name
					cl.tx <- LobbyChoices{lobbiesList()}
				} else {
					unnamed.M[cid].tx <- NameReject{}
				}
			}
		}
	}
}
