package main

import (
	"log"
	"math/rand"
)

// Game is a group of clients playing a game together.
type Game struct {
	tx chan<- CIDClient // from runLobbyMap to this
}

type nMembers struct {
	total  int
	resWin int
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
	names map[CID]string,
	tx chan<- SrvMsg,
	q <-chan struct{},
) Game {
	log.Println("NewGame", gid)
	// see NewLobby.
	rxLobbyMap := make(chan CIDClient)
	go runGame(gid, clients, names, tx, rxLobbyMap, q)
	return Game{rxLobbyMap}
}

func runGame(
	gid GID,
	clients *ClientMap,
	names map[CID]string,
	tx chan<- SrvMsg,
	rx <-chan CIDClient,
	q <-chan struct{},
) {
	// whenever sending on tx, must also select on rx and q to prevent deadlock.

	// all the cids, in a stable order. even if a client disconnects/reconnects,
	// this remains unchanged.
	cids := make([]CID, 0, len(clients.M))
	for cid := range clients.M {
		cids = append(cids, cid)
	}

	// len(cids)/s clients will be spies.
	const s = 4

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
	state := memberChoosing

	// invariant: the client with CID cids[captain] is the current captain.
	captain := 0

	// invariant:
	// state == missionVoting || state == memberVoting <=> members != nil.
	var members []CID

	// invariant: members == nil <=> votes == nil.
	var votes map[CID]bool

	// invariant: 0 <= resPts <= MaxPts.
	resPts := 0

	// invariant: 0 <= spyPts <= MaxPts.
	spyPts := 0

	// invariant: 0 := skip <= MaxSkip.
	skip := 0

	// getNMembers returns the number of clients on this mission, and the number
	// of clients required to vote to pass the mission for this mission to pass.
	// TODO improve
	getNMembers := func() nMembers {
		const m = 5
		ret := len(cids) / m
		return nMembers{ret, (ret / 2) + 1}
	}

	// newMemberChoosing moves the state into memberChoosing and update the
	// appropriate other bits of state.
	newMemberChoosing := func() {
		state = memberChoosing
		members = nil
		votes = nil
		captain++
		if captain == len(cids) {
			captain = 0
		}
	}

	// getCurrentGame returns the current game state in a format readable by
	// clients.
	getCurrentGame := func(cid CID) CurrentGame {
		return CurrentGame{
			IsSpy:      isSpy[cid],
			ResPts:     resPts,
			SpyPts:     spyPts,
			Captain:    cids[captain],
			NumMembers: getNMembers().total,
			Members:    members,
			Active:     state == missionVoting,
		}
	}

	// reconnect tries to reconnect a CIDClient to this game. if it fails, it
	// closes that CIDClient.
	reconnect := func(cl CIDClient) {
		_, ok := clients.M[cl.CID]
		if ok || !hasCID(cids, cl.CID) {
			cl.Close()
		} else {
			clients.Add(cl.CID, cl.Client)
			cl.tx <- getCurrentGame(cl.CID)
		}
	}

	// broadcast broadcasts the current game state to all clients.
	broadcast := func() {
		for cid, cl := range clients.M {
			cl.tx <- getCurrentGame(cid)
		}
	}

	broadcast()

	for {
		select {
		case <-q:
			clients.CloseAll()
			return
		case cl := <-rx:
			reconnect(cl)
		case m := <-clients.C:
			cid := m.CID
			switch ts := m.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				// TODO only do this after a timeout?
				if len(clients.M) == 0 {
					goto out
				}
			case MemberChoose:
				if state != memberChoosing ||
					cid != cids[captain] ||
					len(ts.Members) != getNMembers().total {
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
				if len(votes) != getNMembers().total {
					continue
				}
				success := numTrue(votes) >= getNMembers().resWin
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
	clients.DisconnectAll()
	for {
		select {
		case <-q:
			clients.CloseAll()
			return
		case cl := <-rx:
			_, ok := clients.M[cl.CID]
			if ok || !hasCID(cids, cl.CID) {
				cl.Close()
			} else {
				clients.AddNoSend(cl.CID, cl.Client)
			}
		case tx <- GameClose{gid, clients, names, EndGame{resPts, spyPts, nil}}:
			return
		}
	}
}
