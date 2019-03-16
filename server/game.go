package main

import (
	"log"
	"math/rand"
)

type state uint

const (
	memberChoosing state = iota
	memberVoting
	missionVoting
	gameOver
)

func numTrue(xs map[CID]bool) int {
	ret := 0
	for _, x := range xs {
		if x {
			ret++
		}
	}
	return ret
}

func hasCID(xs []CID, x CID) bool {
	for _, y := range xs {
		if x == y {
			return true
		}
	}
	return false
}

// TODO ensure that at least one spy is included in the first mission?
func runGame(gid GID, tx chan<- LobbyMsg, clients *ClientMap) {
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	cs := clients.ToList()
	// as per lobby.go, MinN <= n <= MaxN
	n := len(cs)
	// n/s clients will be spies
	const s = 4
	// n/m clients each round will be part of a mission
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

	state := memberChoosing
	captainIdx := n - 1
	newMission := func() NewMission {
		captainIdx++
		if captainIdx == n {
			captainIdx = 0
		}
		return NewMission{cs[captainIdx].CID, nMission}
	}
	nm := newMission()
	for i, cl := range cs {
		cl.tx <- FirstMission{isSpy[i], nm}
	}

	// invariant: 0 <= spyWinN, resWinN <= MaxWin
	spyWinN := 0
	resWinN := 0

	// invariant: state == missionVoting <=> members != nil
	var members []CID
	// used for both voting on mission members and voting on mission itself
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
			state = memberVoting
			members = ts.Members
			votes = make(map[CID]bool)
			for _, cl := range cs {
				cl.tx <- MemberPropose{ts.Members}
			}
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
			if numTrue(votes) > n/2 {
				state = missionVoting
				for _, cl := range cs {
					cl.tx <- MemberAccept{}
				}
			} else {
				state = memberChoosing
				msg := newMission()
				for _, cl := range cs {
					cl.tx <- msg
				}
			}
		case MissionVote:
			if state != missionVoting || !hasCID(members, cid) {
				continue
			}
			_, ok := votes[cid]
			if ok {
				continue
			}
			votes[cid] = ts.Vote
			if len(votes) != nMission {
				continue
			}
			success := numTrue(votes) > nMission/2
			if success {
				resWinN++
			} else {
				spyWinN++
			}
			msg := MissionResult{Success: success}
			if resWinN < MaxWin && spyWinN < MaxWin {
				state = missionVoting
				msg.NewMission = newMission()
			} else {
				state = gameOver
			}
			for _, cl := range cs {
				cl.tx <- msg
			}
		}
	}
}
