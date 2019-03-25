package main

import (
	"regexp"
)

var validNameRE = regexp.MustCompile(`[^\s]`)

func validName(s string) bool {
	const maxLen = 32
	return s != "" && len(s) <= maxLen && validNameRE.Match([]byte(s))
}

func runWelcomer(tx chan<- ToLobbyMap, rx <-chan *Client, q <-chan struct{}) {
	clients := NewClientMap()
	for {
		select {
		case <-q:
			return
		case cl := <-rx:
			clients.Add(cl)
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case Connect:
				clients.M[cid].tx <- SetMe{cid}
			case Reconnect:
				cl := clients.Rm(cid)
				cl.CID = ts.Me
				tx <- ClientReconnect{cl, ts.GID}
			case NameChoose:
				if validName(ts.Name) {
					cl := clients.Rm(cid)
					cl.Name = ts.Name
					tx <- ClientAdd{cl}
				} else {
					clients.M[cid].tx <- NameReject{}
				}
			}
		}
	}
}
