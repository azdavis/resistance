package main

func runWelcomer(tx chan<- ToLobbyMap, rx <-chan Client, q <-chan struct{}) {
	clients := NewClientMap()
	next := CID(1)
	for {
		select {
		case <-q:
			clients.CloseAll()
			return
		case cl := <-rx:
			clients.Add(next, cl)
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
				tx <- ClientReconnect{ts.Me, cl, ts.GID}
			case NameChoose:
				if ValidName(ts.Name) {
					tx <- ClientAdd{cid, clients.Rm(cid), ts.Name}
				} else {
					clients.M[cid].tx <- NameReject{}
				}
			}
		}
	}
}
