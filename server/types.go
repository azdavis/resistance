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

// ToServer ////////////////////////////////////////////////////////////////////

// ToServer is a request from the client to change state. The client "requests"
// a Close by closing itself.
type ToServer interface {
	isToServer()
}

func (Close) isToServer()        {}
func (Connect) isToServer()      {}
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

// NameChoose is a request to choose one's name.
type NameChoose struct {
	Name string // desired name
}

// LobbyChoose is a request to choose one's lobby.
type LobbyChoose struct {
	GID // desired GID
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
	Members []CID // the proposed members of the mission
}

// MemberVote is sent by a client to vote for the proposed mission members.
type MemberVote struct {
	Vote bool // whether the client approved of the proposed members
}

// MissionVote is sent by a client to vote on whether the mission should succeed
// or fail.
type MissionVote struct {
	Vote bool // whether the client wants the mission to succeed
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

func (SetMe) isToClient()         {}
func (NameReject) isToClient()    {}
func (LobbyChoices) isToClient()  {}
func (CurrentLobby) isToClient()  {}
func (FirstMission) isToClient()  {}
func (MemberPropose) isToClient() {}
func (MemberAccept) isToClient()  {}
func (MemberReject) isToClient()  {}
func (MissionResult) isToClient() {}

// SetMe sets the client's own CID.
type SetMe struct {
	Me CID
}

// NameReject is sent to a client that requested a name change with
// NameChoose.
type NameReject struct{}

// LobbyChoices is sent to a client who is choosing their lobby.
type LobbyChoices struct {
	Lobbies []Lobby // available lobbies to join
}

// CurrentLobby is sent to a client who is in a lobby whose game has not yet
// started.
type CurrentLobby struct {
	GID               // the GID of this lobby
	Leader  CID       // the leader of this lobby
	Clients []*Client // info about other clients in this lobby
}

// FirstMission is sent to start the game.
type FirstMission struct {
	IsSpy   bool // whether the client is a spy
	Captain CID  // captain of this mission
	Members int  // number of members on this mission
}

// MemberPropose notifies the client that the captain has selected mission
// candidates, and voting on whether the mission will proceed can begin.
type MemberPropose struct {
	Members []CID // CIDs selected by the captain
}

// MemberAccept notifies the client that the proposed members have been
// accepted.
type MemberAccept struct{}

// MemberReject notifies the client that the proposed members have been
// rejected.
type MemberReject struct {
	Captain CID  // captain of new mission
	Members int  // number of members on new mission
	SpyWin  bool // whether the spies get a point
}

// MissionResult notifies the client that voting on the mission has concluded.
type MissionResult struct {
	Success bool // whether this mission succeeded
	Captain CID  // captain of new mission
	Members int  // number of members on new mission
}

// ToLobbyMap //////////////////////////////////////////////////////////////////

// ToLobbyMap is a message to the lobby map.
type ToLobbyMap interface {
	isToLobbyMap()
}

func (ClientAdd) isToLobbyMap()  {}
func (LobbyClose) isToLobbyMap() {}
func (GameCreate) isToLobbyMap() {}
func (GameClose) isToLobbyMap()  {}

// ClientAdd signals that a client is being added to the lobby map.
type ClientAdd struct {
	*Client // the client that is being added
}

// LobbyClose signals that a lobby is closing.
type LobbyClose struct {
	GID               // gid of this lobby
	Clients []*Client // clients coming from this lobby
}

// GameCreate signals that a lobby is turning into a game.
type GameCreate struct {
	GID                // gid of this lobby
	Clients *ClientMap // clients in this lobby
}

// GameClose signals that a game is closing.
type GameClose struct {
	GID               // gid of this game
	Clients []*Client // clients coming from this game
}

// Helper functions ////////////////////////////////////////////////////////////

// tagMsg is a JSON-encoded message.
type tagMsg struct {
	T string          // type of thing to try to parse
	P json.RawMessage // json encoding of thing
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
func (x FirstMission) MarshalJSON() ([]byte, error) {
	type alias FirstMission
	return fromTagMsg("FirstMission", alias(x))
}

// MarshalJSON makes JSON.
func (x MemberPropose) MarshalJSON() ([]byte, error) {
	type alias MemberPropose
	return fromTagMsg("MemberPropose", alias(x))
}

// MarshalJSON makes JSON.
func (x MemberAccept) MarshalJSON() ([]byte, error) {
	type alias MemberAccept
	return fromTagMsg("MemberAccept", alias(x))
}

// MarshalJSON makes JSON.
func (x MemberReject) MarshalJSON() ([]byte, error) {
	type alias MemberReject
	return fromTagMsg("MemberReject", alias(x))
}

// MarshalJSON makes JSON.
func (x MissionResult) MarshalJSON() ([]byte, error) {
	type alias MissionResult
	return fromTagMsg("MissionResult", alias(x))
}
