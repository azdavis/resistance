package main

import (
	"log"
	"math/rand"
)

func runGame(gid GID, leaderID CID, tx chan<- LobbyMsg, clients *ClientMap) {
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	cs := clients.ToList()
	// as per lobby.go, n >= MinN.
	n := len(cs)
	// n/s clients will be spies.
	const s = 4
	// n/m clients each round will be part of a mission.
	const m = 5
	nMission := n / m

	isSpy := make(map[CID]bool)
	for i := n / s; i > 0; /* intentionally empty */ {
		cid := cs[rand.Intn(n)].CID
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
	msg := NewMission{cs[captainIdx].CID, nMission}
	for _, cl := range cs {
		cl.tx <- msg
	}

	for ac := range clients.C {
		cid := ac.CID
		switch ts := ac.ToServer.(type) {
		case Close:
			clients.Rm(cid).Close()
			tx <- LobbyMsg{gid, false, clients.Clear()}
			return
		case MemberChoose:
			if cid != cs[captainIdx].CID || len(ts.Members) != nMission {
				continue
			}
			msg := MemberPropose{ts.Members}
			for _, cl := range cs {
				cl.tx <- msg
			}
		}
	}
}
