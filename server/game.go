package main

import (
	"log"
	"math/rand"
)

type state uint

const (
	memberChoosing state = iota
	memberVoting
	missionRunning
)

// TODO ensure that at least one spy is included in the first mission?
func runGame(gid GID, tx chan<- LobbyMsg, clients *ClientMap) {
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

	// start all false
	isSpy := make([]bool, n)
	for i := n / s; i > 0; /* intentionally empty */ {
		j := rand.Intn(n)
		if isSpy[j] {
			continue
		}
		isSpy[j] = true
		i--
	}
	log.Println("spies", gid, isSpy)

	for i, cl := range cs {
		cl.tx <- SetIsSpy{isSpy[i]}
	}

	var state state
	captainIdx := n - 1
	newMission := func() {
		state = memberChoosing
		captainIdx++
		if captainIdx == n {
			captainIdx = 0
		}
		msg := NewMission{cs[captainIdx].CID, nMission}
		for _, cl := range cs {
			cl.tx <- msg
		}
	}
	newMission()

	votes := make(map[CID]bool)

	for ac := range clients.C {
		cid := ac.CID
		switch ts := ac.ToServer.(type) {
		case Close:
			clients.Rm(cid).Close()
			tx <- LobbyMsg{gid, false, clients.Clear()}
			return
		case MemberChoose:
			if state != memberChoosing ||
				cid != cs[captainIdx].CID ||
				len(ts.Members) != nMission {
				continue
			}
			for _, cl := range cs {
				cl.tx <- MemberPropose{ts.Members}
			}
			state = memberVoting
		case MemberVote:
			if state != memberVoting {
				continue
			}
			_, ok := votes[cid]
			if ok {
				continue
			}
			votes[cid] = ts.Vote
			if len(votes) != n {
				continue
			}
			yes := 0
			for _, v := range votes {
				if v {
					yes++
				}
			}
			if yes > n/2 {
				state = missionRunning
				for _, cl := range cs {
					cl.tx <- MemberAccept{}
				}
			} else {
				newMission()
			}
		}
	}
}
