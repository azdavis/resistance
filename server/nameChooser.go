package main

func runNameChooser(tx chan<- *Client, rx <-chan *Client) {
	clients := NewClientMap()
	for {
		select {
		case cl := <-rx:
			clients.Add(cl)
		case ac := <-clients.C:
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
			case NameChoose:
				if validName(ts.Name) {
					cl := clients.Rm(cid)
					cl.name = ts.Name
					tx <- cl
				} else {
					clients.M[cid].tx <- NameChoosing{Valid: false}
				}
			}
		}
	}
}
