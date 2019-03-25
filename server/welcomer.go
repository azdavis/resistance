package main

import (
	"log"
	"regexp"
)

var validNameRE = regexp.MustCompile(`[^\s]`)

func validName(s string) bool {
	const maxLen = 32
	return s != "" && len(s) <= maxLen && validNameRE.Match([]byte(s))
}

func runWelcomer(tx chan<- ToLobbyMap, rx <-chan *Client, q <-chan struct{}) {
	log.Println("enter runWelcomer")
	defer log.Println("exit runWelcomer")
	clients := NewClientMap()
	next := CID(1)
	for {
		select {
		case <-q:
			clients.KillAll()
			return
		case cl := <-rx:
			cl.CID = next
			next++
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
