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

// NewGame returns a new Game.
func NewGame(gid GID, clients *ClientMap, tx chan<- ToLobbyMap) Game {
	// see NewLobby.
	rxLobbyMap := make(chan *Client)
	g := Game{
		GID: gid,
		tx:  rxLobbyMap,
	}
	go runGame(gid, clients, tx, rxLobbyMap)
	return g
}

// TODO improve numbers for mission size / fails required to fail mission?
func runGame(
	gid GID,
	clients *ClientMap,
	tx chan<- ToLobbyMap,
	rx <-chan *Client,
) {
	// whenever sending on tx, must also select with rx to prevent deadlock.
	log.Println("enter runGame", gid)
	defer log.Println("exit runGame", gid)

	reconnect := func(cl *Client) {
		_, ok := clients.M[cl.CID]
		if ok {
			// oof
			cl.Kill()
		} else {
			clients.Add(cl)
			// TODO send a message to cl to get it up-to-date
		}
	}

	// all the cids, in a stable order.
	cids := make([]CID, 0, len(clients.M))
	for cid := range clients.M {
		cids = append(cids, cid)
	}

	// len(cids)/s clients will be spies.
	const s = 4

	// len(cids)/m clients each round will be part of a mission.
	const m = 5
	nMission := len(cids) / m

	// invariant: isSpy[i] <=> the client with CID cids[i] is a spy.
	isSpy := make([]bool, len(cids))
	for i := len(cids) / s; i > 0; /* intentionally empty */ {
		j := rand.Intn(len(cids))
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

	// invariant: the client with CID cids[captain] is the current captain.
	captain := 0

	// update captain.
	nextCaptain := func() {
		state = memberChoosing
		captain++
		if captain == len(cids) {
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

	msg := FirstMission{Captain: cids[captain], Members: nMission}
	for i, cid := range cids {
		msg.IsSpy = isSpy[i]
		clients.M[cid].tx <- msg
	}

	for {
		select {
		case cl := <-rx:
			reconnect(cl)
		case ac := <-clients.C:
			log.Printf("runGame %v ac: %+v", gid, ac)
			cid := ac.CID
			switch ts := ac.ToServer.(type) {
			case Close:
				clients.Rm(cid).Close()
				// TODO only do this after a timeout?
				for len(clients.M) == 0 {
					select {
					case tx <- GameClose{gid, []*Client{}}:
						return
					case cl := <-rx:
						reconnect(cl)
					}
				}
			case MemberChoose:
				if state != memberChoosing ||
					cid != cids[captain] ||
					len(ts.Members) != nMission {
					continue
				}
				state = memberVoting
				votes = make(map[CID]bool)
				members = ts.Members
				for _, cl := range clients.M {
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
				if len(votes) != len(cids) {
					continue
				}
				if numTrue(votes) > len(votes)/2 {
					state = missionVoting
					votes = make(map[CID]bool)
					for _, cl := range clients.M {
						cl.tx <- MemberAccept{}
					}
				} else {
					nextCaptain()
					skip++
					spyDidWin := skip == MaxSkip
					msg := MemberReject{
						Captain: cids[captain],
						Members: nMission,
						SpyWin:  spyDidWin,
					}
					for _, cl := range clients.M {
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
					msg.Captain = cids[captain]
					msg.Members = nMission
				} else {
					state = gameOver
				}
				for _, cl := range clients.M {
					cl.tx <- msg
				}
			case GameLeave:
				if state != gameOver {
					continue
				}
				cl := clients.Rm(cid)
				for len(clients.M) == 0 {
					select {
					case tx <- GameClose{gid, []*Client{cl}}:
						return
					case cl := <-rx:
						reconnect(cl)
					}
				}
				select {
				case tx <- ClientAdd{cl}:
				case cl := <-rx:
					reconnect(cl)
				}
			}
		}
	}
}
