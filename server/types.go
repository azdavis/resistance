package main

import (
	"encoding/json"
	"fmt"
)

// Helper types ////////////////////////////////////////////////////////////////

// CID is a unique identifier for a Client.
type CID uint64

// GID is a unique identifier for a Game (or a Lobby, which will become a Game).
type GID uint64

// Action is a CID + ToServer.
type Action struct {
	CID
	ToServer
}

// LobbyMsg is sent from a lobby or game to the lobby manager.
type LobbyMsg struct {
	GID               // gid of this lobby
	Close   bool      // whether to close this lobby
	Clients []*Client // clients coming form this lobby
}

// ToServer ////////////////////////////////////////////////////////////////////

// ToServer is a request from the client to change state. The client "requests"
// a Close by closing itself.
//
// These should be kept in sync with types.ts.
type ToServer interface {
	isToServer()
}

func (Close) isToServer()         {}
func (NameChoose) isToServer()    {}
func (LobbyChoose) isToServer()   {}
func (LobbyLeave) isToServer()    {}
func (LobbyCreate) isToServer()   {}
func (GameStart) isToServer()     {}
func (MissionChoose) isToServer() {}

// Close means the client closed itself. No further Actions will follow from
// this client.
type Close struct{}

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

// MissionChoose is a request from the mission captain to use the given CIDs
// as the members of this mission.
type MissionChoose struct {
	Members []CID
}

// ToClient ////////////////////////////////////////////////////////////////////

// ToClient is sent to the client to change the client's state. It may be sent
// in direct reply to a client's ToServer, or it may be sent because the client
// was transitively affected by another client's ToServer.
//
// These should be kept in sync with types.ts.
type ToClient interface {
	json.Marshaler
	isToClient()
}

func (RejectName) isToClient()   {}
func (LobbyChoices) isToClient() {}
func (CurrentLobby) isToClient() {}
func (SetIsSpy) isToClient()     {}
func (NewMission) isToClient()   {}

// RejectName is sent to a client that requested a name change with
// NameChoose.
type RejectName struct{}

// LobbyChoices is sent to a client who is choosing their lobby.
type LobbyChoices struct {
	Lobbies []Lobby // available lobbies to join
}

// CurrentLobby is sent to a client who is in a lobby whose game has not yet
// started.
type CurrentLobby struct {
	Me      CID       // the client's own CID
	Leader  CID       // info about this lobby
	Clients []*Client // info about other clients in this lobby
}

// SetIsSpy sets whether the clients is a spy or not.
type SetIsSpy struct {
	IsSpy bool // if false, client is resistance
}

// NewMission notifies the client that a new mission has started.
type NewMission struct {
	Captain CID // captain of this mission
	N       int // number of members on this mission
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
	case "MissionChoose":
		var msg MissionChoose
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
func (pc RejectName) MarshalJSON() ([]byte, error) {
	type alias RejectName
	return fromTagMsg("RejectName", alias(pc))
}

// MarshalJSON makes JSON.
func (pc LobbyChoices) MarshalJSON() ([]byte, error) {
	type alias LobbyChoices
	return fromTagMsg("LobbyChoices", alias(pc))
}

// MarshalJSON makes JSON.
func (pc CurrentLobby) MarshalJSON() ([]byte, error) {
	type alias CurrentLobby
	return fromTagMsg("CurrentLobby", alias(pc))
}

// MarshalJSON makes JSON.
func (pc SetIsSpy) MarshalJSON() ([]byte, error) {
	type alias SetIsSpy
	return fromTagMsg("SetIsSpy", alias(pc))
}

// MarshalJSON makes JSON.
func (pc NewMission) MarshalJSON() ([]byte, error) {
	type alias NewMission
	return fromTagMsg("NewMission", alias(pc))
}
