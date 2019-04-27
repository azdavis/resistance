package main

import (
	"encoding/json"
	"fmt"
)

// CIDToServer is a CID + ToServer.
type CIDToServer struct {
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

// Dest represents a place to put Actions, and what CID to tag them with.
type Dest struct {
	CID
	C chan<- CIDToServer
}

// SrvMsg is a message to runServer.
type SrvMsg interface {
	isSrvMsg()
}

func (Client) isSrvMsg()      {}
func (NamedClient) isSrvMsg() {}
func (LobbyClose) isSrvMsg()  {}
func (GameClose) isSrvMsg()   {}

// NamedClient signals that a client with a name is being added to the server.
type NamedClient struct {
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
	type t SetMe
	return fromTagMsg("SetMe", t(x))
}

// MarshalJSON makes JSON.
func (x NameReject) MarshalJSON() ([]byte, error) {
	type t NameReject
	return fromTagMsg("NameReject", t(x))
}

// MarshalJSON makes JSON.
func (x LobbyChoices) MarshalJSON() ([]byte, error) {
	type t LobbyChoices
	return fromTagMsg("LobbyChoices", t(x))
}

// MarshalJSON makes JSON.
func (x CurrentLobby) MarshalJSON() ([]byte, error) {
	type t CurrentLobby
	return fromTagMsg("CurrentLobby", t(x))
}

// MarshalJSON makes JSON.
func (x CurrentGame) MarshalJSON() ([]byte, error) {
	type t CurrentGame
	return fromTagMsg("CurrentGame", t(x))
}

// MarshalJSON makes JSON.
func (x EndGame) MarshalJSON() ([]byte, error) {
	type t EndGame
	return fromTagMsg("EndGame", t(x))
}
