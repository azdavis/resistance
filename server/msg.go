package main

import (
	"encoding/json"
	"fmt"
)

// ToServer ////////////////////////////////////////////////////////////////////

// ToServer is a request from the client to change state. The client "requests"
// a Close by closing itself.
type ToServer interface {
	isToServer()
}

func (Close) isToServer()       {}
func (NameChoose) isToServer()  {}
func (LobbyChoose) isToServer() {}
func (LobbyLeave) isToServer()  {}
func (LobbyCreate) isToServer() {}
func (GameStart) isToServer()   {}

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

// ToClient //////////////////////////////////////////////////////////////////

// ToClient is sent to the client to change the client's state. It may be sent
// in direct reply to a client's ToServer, or it may be sent because the client
// was transitively affected by another client's ToServer.
type ToClient interface {
	json.Marshaler
	isToClient()
}

func (NameChoosing) isToClient()  {}
func (LobbyChoosing) isToClient() {}
func (LobbyWaiting) isToClient()  {}

// NameChoosing is sent to a client that requested a name change with
// NameChoose. Valid is always false, since if name was valid, we would send
// LobbyChoosing instead.
type NameChoosing struct {
	Valid bool // whether the name was valid
}

// LobbyInfo contains info about a Lobby.
type LobbyInfo struct {
	GID
	Leader string
}

// LobbyChoosing is sent to a client who is choosing their lobby.
type LobbyChoosing struct {
	Lobbies []LobbyInfo // available lobbies to join
}

// ClientInfo contains info about a Client.
type ClientInfo struct {
	CID
	Name string
}

// LobbyWaiting is sent to a client who is in a lobby whose game has not yet
// started.
type LobbyWaiting struct {
	Self    CID          // the client's own CID
	Leader  CID          // info about this lobby
	Clients []ClientInfo // info about other clients in this lobby
}

// helpers /////////////////////////////////////////////////////////////////////

// tagMsg is a JSON-encoded message.
type tagMsg struct {
	T string          // type of thing to try to parse
	P json.RawMessage // json encoding of thing
}

// fromTagMsg creates a JSON-encoded tagMsg.
func fromTagMsg(t string, p interface{}) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return json.Marshal(tagMsg{T: t, P: json.RawMessage(bs)})
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
	default:
		return nil, fmt.Errorf("unknown tag: %s", tm.T)
	}
}

// MarshalJSON makes JSON.
func (pc NameChoosing) MarshalJSON() ([]byte, error) {
	type alias NameChoosing
	return fromTagMsg("NameChoosing", alias(pc))
}

// MarshalJSON makes JSON.
func (pc LobbyChoosing) MarshalJSON() ([]byte, error) {
	type alias LobbyChoosing
	return fromTagMsg("LobbyChoosing", alias(pc))
}

// MarshalJSON makes JSON.
func (pc LobbyWaiting) MarshalJSON() ([]byte, error) {
	type alias LobbyWaiting
	return fromTagMsg("LobbyWaiting", alias(pc))
}
