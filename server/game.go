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

func hasCID(xs []CID, y CID) bool {
	for _, x := range xs {
		if y == x {
			return true
		}
	}
	return false
}

// TODO ensure that at least one spy is included in the first mission?
// TODO improve numbers for mission size / fails required to fail mission?
func runGame(gid GID, tx chan<- LobbyMsg, clients *ClientMap) {
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	// all the clients, in a stable order.
	cs := clients.ToList()

	// len(cs)/s clients will be spies.
	const s = 4
	// len(cs)/m clients each round will be part of a mission.
	const m = 5
	nMission := len(cs) / m

	// invariant: isSpy[i] <=> cs[i] is a spy.
	isSpy := make([]bool, len(cs))
	for i := len(cs) / s; i > 0; /* intentionally empty */ {
		j := rand.Intn(len(cs))
		if isSpy[j] {
			continue
		}
		isSpy[j] = true
		i--
	}
	log.Printf("runGame %v spies: %+v", gid, isSpy)

	// current state.
	// invariant: state == gameOver <=> resWin == MaxWin || spyWin == MaxWin
	state := memberChoosing
	// invariant: cs[captain] is the current captain.
	captain := 0
	// update captain.
	nextCaptain := func() {
		state = memberChoosing
		captain++
		if captain == len(cs) {
			captain = 0
		}
	}

	// invariant: 0 <= resWin <= MaxWin
	resWin := 0
	// invariant: 0 <= spyWin <= MaxWin
	spyWin := 0

	// invariant: 0 := skip <= MaxSkip
	skip := 0

	// invariant: state == missionVoting <=> members != nil
	var members []CID

	// used for both voting on mission members and voting on mission itself
	votes := make(map[CID]bool)

	msg := FirstMission{Captain: cs[captain].CID, Members: nMission}
	for i, cl := range cs {
		msg.IsSpy = isSpy[i]
		cl.tx <- msg
	}

	for ac := range clients.C {
		log.Printf("runGame %v ac: %+v", gid, ac)
		cid := ac.CID
		switch ts := ac.ToServer.(type) {
		case Close:
			// TODO allow reconnecting?
			clients.Rm(cid).Close()
			tx <- LobbyMsg{gid, false, clients.Clear()}
			return
		case MemberChoose:
			if state != memberChoosing ||
				cid != cs[captain].CID ||
				len(ts.Members) != nMission {
				continue
			}
			state = memberVoting
			votes = make(map[CID]bool)
			members = ts.Members
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
			if len(votes) != len(cs) {
				continue
			}
			if numTrue(votes) > len(votes)/2 {
				state = missionVoting
				votes = make(map[CID]bool)
				for _, cl := range cs {
					cl.tx <- MemberAccept{}
				}
			} else {
				nextCaptain()
				skip++
				spyDidWin := skip == MaxSkip
				msg := MemberReject{
					Captain: cs[captain].CID,
					Members: nMission,
					SpyWin:  spyDidWin,
				}
				for _, cl := range cs {
					cl.tx <- msg
				}
				if spyDidWin {
					spyWin++
					skip = 0
				}
				if spyWin >= MaxWin {
					state = gameOver
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
			members = nil
			success := numTrue(votes) > nMission/2
			if success {
				resWin++
			} else {
				spyWin++
			}
			msg := MissionResult{Success: success}
			if resWin < MaxWin && spyWin < MaxWin {
				nextCaptain()
				msg.Captain = cs[captain].CID
				msg.Members = nMission
			} else {
				state = gameOver
			}
			for _, cl := range cs {
				cl.tx <- msg
			}
		case GameLeave:
			if state != gameOver {
				continue
			}
			tx <- LobbyMsg{gid, false, []*Client{clients.Rm(cid)}}
			if len(clients.M) == 0 {
				return
			}
		}
	}
}
