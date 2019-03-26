package main

import (
	"regexp"
)

var validNameRE = regexp.MustCompile(`[^\s]`)

func validName(s string) bool {
	const maxLen = 32
	return s != "" && len(s) <= maxLen && validNameRE.Match([]byte(s))
}

func runWelcomer(tx chan<- ToLobbyMap, rx <-chan Client, q <-chan struct{}) {
	clients := NewClientMap()
	next := CID(1)
	for {
		select {
		case <-q:
			clients.CloseAll()
			return
		case cl := <-rx:
			clients.Add(CIDClient{next, cl})
			next++
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case Connect:
				clients.M[cid].tx <- SetMe{cid}
			case Reconnect:
				cl := clients.Rm(cid)
				tx <- ClientReconnect{CIDClient{ts.Me, cl}, ts.GID}
			case NameChoose:
				if validName(ts.Name) {
					tx <- ClientAdd{CIDClient{cid, clients.Rm(cid)}}
				} else {
					clients.M[cid].tx <- NameReject{}
				}
			}
		}
	}
}
