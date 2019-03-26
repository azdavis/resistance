// These should be kept in sync with types.ts.
package main

import (
	"encoding/json"
	"fmt"
)

// Helper types ////////////////////////////////////////////////////////////////

// CID is a unique identifier for a Client. 0 means 'no CID'.
type CID uint64

// GID is a unique identifier for a Game (or a Lobby, which will become a Game).
// 0 means 'no GID'.
type GID uint64

// Action is a CID + ToServer.
type Action struct {
	CID
	ToServer
}

// CIDClient is a CID + Client.
type CIDClient struct {
	CID
	Client
}

// ClientInfo represents info about a Client.
type ClientInfo struct {
	CID
	Name string
}

// ToServer ////////////////////////////////////////////////////////////////////

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

// ToClient ////////////////////////////////////////////////////////////////////

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
// invariant: Me != 0
type SetMe struct {
	Me CID
}

// NameReject is sent to a client that requested a name change with
// NameChoose.
type NameReject struct{}

// LobbyChoices is sent to a client who is choosing their lobby.
// invariant: Lobbies != nil
// invariant: for all x in Lobbies, x.GID != 0
type LobbyChoices struct {
	Lobbies []Lobby
}

// CurrentLobby is sent to a client who is in a lobby whose game has not yet
// started.
// invariant: GID != 0
// invariant: Leader != 0
// invariant: Clients != nil
// invariant: for all x in Clients, x.CID != 0
type CurrentLobby struct {
	GID
	Leader  CID
	Clients []ClientInfo
}

// CurrentGame represents an in-progress game.
// invariant: 0 <= ResPts < MaxPts
// invariant: 0 <= SpyPts < MaxPts
// invariant: Members != nil => len(Members) == NumMembers
// invariant: Active => Members != nil
// invariant: for all x in Members, x != 0
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
// invariant: 0 <= ResPts <= MaxPts
// invariant: 0 <= SpyPts <= MaxPts
// invariant: ResPts == MaxPts <=> SpyPts != MaxPts
// invariant: Lobbies != nil
// invariant: for all x in Lobbies, x.GID != 0
type EndGame struct {
	ResPts  int
	SpyPts  int
	Lobbies []Lobby
}

// SrvMsg //////////////////////////////////////////////////////////////////////

// SrvMsg is a message to runServer.
type SrvMsg interface {
	isSrvMsg()
}

func (Client) isSrvMsg()     {}
func (ClientAdd) isSrvMsg()  {}
func (LobbyClose) isSrvMsg() {}
func (GameClose) isSrvMsg()  {}

// ClientAdd signals that a client is being added to the lobby map.
type ClientAdd struct {
	CID
	Client
	Name string
}

// LobbyClose signals that a lobby is closing.
type LobbyClose struct {
	MakeGame bool
	GID
	Clients *ClientMap
	Names   map[CID]string
}

// GameClose signals that a game is closing.
// invariant: EndGame.Lobbies == nil when the GameClose is received from a game
type GameClose struct {
	GID
	Clients *ClientMap
	Names   map[CID]string
	EndGame
}

// Helper functions ////////////////////////////////////////////////////////////

// tagMsg is a JSON-encoded message.
type tagMsg struct {
	T string
	P json.RawMessage
}

// UnmarshalJSONToServer tries to turn a JSON encoding of a tagMsg into a
// ToServer.
func UnmarshalJSONToServer(bs []byte) (ToServer, error) {
	var tm tagMsg
	err := json.Unmarshal(bs, &tm)
	if err != nil {
		return nil, err
	}
	switch tm.T {
	case "Connect":
		var msg Connect
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "Reconnect":
		var msg Reconnect
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "NameChoose":
		var msg NameChoose
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "LobbyChoose":
		var msg LobbyChoose
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "LobbyLeave":
		var msg LobbyLeave
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "LobbyCreate":
		var msg LobbyCreate
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "GameStart":
		var msg GameStart
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "MemberChoose":
		var msg MemberChoose
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "MemberVote":
		var msg MemberVote
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "MissionVote":
		var msg MissionVote
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "GameLeave":
		var msg GameLeave
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	default:
		return nil, fmt.Errorf("unknown tag: %s", tm.T)
	}
}

// fromTagMsg creates a JSON-encoded tagMsg.
func fromTagMsg(t string, p interface{}) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return json.Marshal(tagMsg{T: t, P: json.RawMessage(bs)})
}

// MarshalJSON makes JSON.
func (x SetMe) MarshalJSON() ([]byte, error) {
	type alias SetMe
	return fromTagMsg("SetMe", alias(x))
}

// MarshalJSON makes JSON.
func (x NameReject) MarshalJSON() ([]byte, error) {
	type alias NameReject
	return fromTagMsg("NameReject", alias(x))
}

// MarshalJSON makes JSON.
func (x LobbyChoices) MarshalJSON() ([]byte, error) {
	type alias LobbyChoices
	return fromTagMsg("LobbyChoices", alias(x))
}

// MarshalJSON makes JSON.
func (x CurrentLobby) MarshalJSON() ([]byte, error) {
	type alias CurrentLobby
	return fromTagMsg("CurrentLobby", alias(x))
}

// MarshalJSON makes JSON.
func (x CurrentGame) MarshalJSON() ([]byte, error) {
	type alias CurrentGame
	return fromTagMsg("CurrentGame", alias(x))
}

// MarshalJSON makes JSON.
func (x EndGame) MarshalJSON() ([]byte, error) {
	type alias EndGame
	return fromTagMsg("EndGame", alias(x))
}
