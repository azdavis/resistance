// This should be kept in sync with client/src/shared.ts.
package main

import (
	"encoding/json"
)

// MinN is the minimum number of players in a game.
const MinN = 5

// MaxN is the maximum number of players in a game.
const MaxN = 7

// OkGameSize returns whether n is an acceptable number of players for the game.
func OkGameSize(n int) bool { return MinN <= n && n <= MaxN }

// MaxPts is the number of wins either side must accumulate before the game is
// over.
const MaxPts = 3

// MaxSkip is the number of missions that can be skipped in a row before the
// spies automatically get a point.
const MaxSkip = 3

// CID is a unique identifier for a Client. 0 means 'no CID'.
type CID uint64

// GID is a unique identifier for a Game (or a Lobby, which will become a Game).
// 0 means 'no GID'.
type GID uint64

// Lobby represents a group of clients all waiting for the same game to
// start.
type Lobby struct {
	GID                       // unique
	Leader string             // leader name
	tx     chan<- NamedClient // from runLobbyMap to this
}

// Client is a player of the game.
type Client struct {
	tx chan ToClient   // orders for the client
	rx chan<- ToServer // requests from the client
	d  chan<- Dest     // what to update the ultimate destination of rx to
	q  chan struct{}   // close on Close
}

// ToServer is a request from the client to change state. The client "requests"
// a Close by closing itself.
type ToServer interface {
	isToServer()
}

func (Close) isToServer()        {}
func (Connect) isToServer()      {}
func (Reconnect) isToServer()    {}
func (NameChoose) isToServer()   {}
func (LobbyChoose) isToServer()  {}
func (LobbyLeave) isToServer()   {}
func (LobbyCreate) isToServer()  {}
func (GameStart) isToServer()    {}
func (MemberChoose) isToServer() {}
func (MemberVote) isToServer()   {}
func (MissionVote) isToServer()  {}
func (GameLeave) isToServer()    {}

// Close means the client closed itself. No further Actions will follow from
// this client.
type Close struct{}

// Connect means the client just connected.
type Connect struct{}

// Reconnect means the client just reconnected. TODO it's not secure to allow
// the client to send this without some kind of cryptographic signature. With
// the current setup, once a client disconnects, any other client can pretend to
// be that client and the server will be none the wiser.
type Reconnect struct {
	Me CID
	GID
}

// NameChoose is a request to choose one's name.
type NameChoose struct {
	Name string
}

// LobbyChoose is a request to choose one's lobby.
type LobbyChoose struct {
	GID
}

// LobbyLeave is a request to leave the client's current lobby.
type LobbyLeave struct{}

// LobbyCreate is a request to create a new lobby, with oneself as the leader.
type LobbyCreate struct{}

// GameStart is a request to start the game.
type GameStart struct{}

// MemberChoose is a request from the captain to use the given CIDs as the
// members of this mission.
type MemberChoose struct {
	Members []CID
}

// MemberVote is sent by a client to vote for the proposed mission members.
type MemberVote struct {
	Vote bool
}

// MissionVote is sent by a client to vote on whether the mission should succeed
// or fail.
type MissionVote struct {
	Vote bool
}

// GameLeave is sent by a client who is leaving the game they are in.
type GameLeave struct{}

// ToClient is sent to the client to change the client's state. It may be sent
// in direct reply to a client's ToServer, or it may be sent because the client
// was transitively affected by another client's ToServer.
type ToClient interface {
	json.Marshaler
	isToClient()
}

func (SetMe) isToClient()        {}
func (NameReject) isToClient()   {}
func (LobbyChoices) isToClient() {}
func (CurrentLobby) isToClient() {}
func (CurrentGame) isToClient()  {}
func (EndGame) isToClient()      {}

// SetMe sets the client's own CID.
// invariant: Me != 0.
type SetMe struct {
	Me CID
}

// NameReject is sent to a client that requested a name change with
// NameChoose.
type NameReject struct{}

// LobbyChoices is sent to a client who is choosing their lobby.
// invariant: Lobbies != nil.
// invariant: for all x in Lobbies, x.GID != 0.
type LobbyChoices struct {
	Lobbies []Lobby
}

// CurrentLobby is sent to a client who is in a lobby whose game has not yet
// started.
// invariant: GID != 0.
// invariant: Leader != 0.
// invariant: Clients != nil.
// invariant: for all x in Clients, x.CID != 0.
type CurrentLobby struct {
	GID
	Leader  CID
	Clients []ClientInfo
}

// CurrentGame represents an in-progress game.
// invariant: 0 <= ResPts < MaxPts.
// invariant: 0 <= SpyPts < MaxPts.
// invariant: Members != nil => len(Members) == NumMembers.
// invariant: Active => Members != nil.
// invariant: for all x in Members, x != 0.
type CurrentGame struct {
	IsSpy      bool
	ResPts     int
	SpyPts     int
	Captain    CID
	NumMembers int
	Members    []CID
	Active     bool
}

// EndGame represents an ended game.
// invariant: 0 <= ResPts <= MaxPts.
// invariant: 0 <= SpyPts <= MaxPts.
// invariant: ResPts == MaxPts <=> SpyPts != MaxPts.
// invariant: Lobbies != nil.
// invariant: for all x in Lobbies, x.GID != 0.
type EndGame struct {
	ResPts  int
	SpyPts  int
	Lobbies []Lobby
}
