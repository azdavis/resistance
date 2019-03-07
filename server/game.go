package main

import (
	"log"
	"math/rand"
)

func runGame(gid GID, leaderID CID, tx chan<- LobbyMsg, clients *ClientMap) {
	// about 1/s of the clients will be spies.
	const s = 4
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	cs := clients.ToList()

	isSpy := make(map[CID]bool)
	for i := len(clients.M) / s; i > 0; /* intentionally empty */ {
		cid := cs[rand.Intn(len(cs))].CID
		if !isSpy[cid] {
			isSpy[cid] = true
			i--
		}
	}
	log.Println("spies", gid, isSpy)

	for _, cl := range cs {
		cl.tx <- SetIsSpy{isSpy[cl.CID]}
	}

	captainIdx := 0
	for i, cl := range cs {
		cl.tx <- NewMission{captainIdx == i}
	}

	for {
		ac := <-clients.C
		cid := ac.CID
		switch ac.ToServer.(type) {
		case Close:
			clients.Rm(cid).Close()
			tx <- LobbyMsg{gid, false, cs}
			return
		}
	}
}
