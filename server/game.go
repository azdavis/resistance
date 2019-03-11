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
	captainIdx := 0
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
			msg := MemberPropose{ts.Members}
			for _, cl := range cs {
				cl.tx <- msg
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
			start := yes > n/2
			msg := MemberResult{start}
			for _, cl := range cs {
				cl.tx <- msg
			}
			if start {
				state = missionRunning
			} else {
				newMission()
			}
		}
	}
}
