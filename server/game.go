package main

import (
	"log"
	"math/rand"
)

// Game is a group of clients playing a game together.
type Game struct {
	GID                // unique
	tx  chan<- *Client // from runLobbyMap to this
}

type state uint

const (
	memberChoosing state = iota
	memberVoting
	missionVoting
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

// NewGame returns a new Game.
func NewGame(
	gid GID,
	clients *ClientMap,
	tx chan<- ToLobbyMap,
	q <-chan struct{},
) Game {
	// see NewLobby.
	rxLobbyMap := make(chan *Client)
	g := Game{
		GID: gid,
		tx:  rxLobbyMap,
	}
	go runGame(gid, clients, tx, rxLobbyMap, q)
	return g
}

// TODO improve numbers for mission size / fails required to fail mission?
func runGame(
	gid GID,
	clients *ClientMap,
	tx chan<- ToLobbyMap,
	rx <-chan *Client,
	q <-chan struct{},
) {
	// whenever sending on tx, must also select on rx and q to prevent deadlock.
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	// all the cids, in a stable order.
	cids := make([]CID, 0, len(clients.M))
	for cid := range clients.M {
		cids = append(cids, cid)
	}

	// all the names.
	names := make(map[CID]string)
	for cid, cl := range clients.M {
		names[cid] = cl.Name
	}

	// len(cids)/s clients will be spies.
	const s = 4

	// len(cids)/m clients each round will be part of a mission.
	const m = 5
	nMission := len(cids) / m

	// invariant: isSpy[cid] <=> the client with CID cid is a spy.
	isSpy := make(map[CID]bool, len(cids))
	for i := len(cids) / s; i > 0; /* intentionally empty */ {
		cid := cids[rand.Intn(len(cids))]
		if isSpy[cid] {
			continue
		}
		isSpy[cid] = true
		i--
	}

	// current state.
	// invariant: state == gameOver <=> resPts == MaxPts || spyPts == MaxPts
	state := memberChoosing

	// invariant: the client with CID cids[captain] is the current captain.
	captain := 0

	// invariant:
	// state == missionVoting || state == memberVoting <=> members != nil
	var members []CID

	// invariant: members == nil <=> votes == nil
	var votes map[CID]bool

	// invariant: 0 <= resPts <= MaxPts
	resPts := 0

	// invariant: 0 <= spyPts <= MaxPts
	spyPts := 0

	// invariant: 0 := skip <= MaxSkip
	skip := 0

	newMemberChoosing := func() {
		state = memberChoosing
		members = nil
		votes = nil
		captain++
		if captain == len(cids) {
			captain = 0
		}
	}

	getCurrentGame := func(cid CID) CurrentGame {
		return CurrentGame{
			IsSpy:      isSpy[cid],
			ResPts:     resPts,
			SpyPts:     spyPts,
			Captain:    cids[captain],
			NumMembers: nMission,
			Members:    members,
			Active:     state == missionVoting,
		}
	}

	reconnect := func(cl *Client) {
		_, ok := clients.M[cl.CID]
		if ok || !hasCID(cids, cl.CID) {
			cl.Kill()
		} else {
			clients.Add(cl)
			cl.Name = names[cl.CID]
			cl.tx <- getCurrentGame(cl.CID)
		}
	}

	broadcast := func() {
		for _, cl := range clients.M {
			cl.tx <- getCurrentGame(cl.CID)
		}
	}

	broadcast()
	log.Printf("runGame gid: %v, names: %+v, isSpy: %+v", gid, names, isSpy)

	for {
		select {
		case <-q:
			clients.KillAll()
			return
		case cl := <-rx:
			reconnect(cl)
		case ac := <-clients.C:
			log.Printf("runGame %v ac: %+v", gid, ac)
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				// TODO only do this after a timeout?
				if len(clients.M) == 0 {
					goto out
				}
			case MemberChoose:
				if state != memberChoosing ||
					cid != cids[captain] ||
					len(ts.Members) != nMission {
					continue
				}
				state = memberVoting
				members = ts.Members
				votes = make(map[CID]bool)
				broadcast()
			case MemberVote:
				if state != memberVoting {
					continue
				}
				_, ok := votes[cid]
				if ok {
					continue
				}
				votes[cid] = ts.Vote
				if len(votes) != len(cids) {
					continue
				}
				if numTrue(votes) > len(votes)/2 {
					state = missionVoting
					votes = make(map[CID]bool)
				} else {
					newMemberChoosing()
					skip++
					spyGetPt := skip == MaxSkip
					if spyGetPt {
						spyPts++
						skip = 0
					}
					if spyPts >= MaxPts {
						goto out
					}
				}
				broadcast()
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
					resPts++
				} else {
					spyPts++
				}
				if resPts >= MaxPts || spyPts >= MaxPts {
					goto out
				}
				newMemberChoosing()
				broadcast()
			}
		}
	}

out:
	cs := clients.Clear()
	for {
		select {
		case <-q:
			for _, cl := range cs {
				cl.Kill()
			}
			return
		case cl := <-rx:
			cs = append(cs, cl)
		case tx <- GameClose{gid, cs, EndGame{resPts, spyPts, nil}}:
			return
		}
	}
}
