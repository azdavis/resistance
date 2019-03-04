package main

import (
	"encoding/json"
	"fmt"
)

// ToServer ////////////////////////////////////////////////////////////////////

// ToServer is a usually parsed tagMsg.P sent from a client. It is a request
// from the client to change state. The only ToServer that does not originate
// from a tagMsg is Close. The client "sends" a Close ToServer by closing
// itself.
type ToServer interface {
	isToServer()
}

func (Close) isToServer()       {}
func (NameChoose) isToServer()  {}
func (PartyChoose) isToServer() {}
func (PartyLeave) isToServer()  {}
func (PartyCreate) isToServer() {}
func (GameStart) isToServer()   {}

// Close means the client closed itself. No further Actions will follow from
// this client.
type Close struct{}

// NameChoose is a request to choose one's name.
type NameChoose struct {
	Name string // desired name
}

// PartyChoose is a request to choose one's party.
type PartyChoose struct {
	PID // desired PID
}

// PartyLeave is a request to leave the client's current party.
type PartyLeave struct{}

// PartyCreate is a request to create a new party, with oneself as the leader.
type PartyCreate struct{}

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
func (PartyChoosing) isToClient() {}
func (PartyWaiting) isToClient()  {}

// NameChoosing is sent to a client that requested a name change with
// NameChoose. Valid is always false, since if name was valid, we would send
// PartyChoosing instead.
type NameChoosing struct {
	Valid bool // whether the name was valid
}

// PartyInfo contains info about a Party.
type PartyInfo struct {
	PID
	Leader string
}

// PartyChoosing is sent to a client who is choosing their party.
type PartyChoosing struct {
	Parties []PartyInfo // available parties to join
}

// ClientInfo contains info about a Client.
type ClientInfo struct {
	CID
	Name string
}

// PartyWaiting is sent to a client who is in a party whose game has not yet
// started.
type PartyWaiting struct {
	Self    CID          // the client's own CID
	Leader  CID          // info about this party
	Clients []ClientInfo // info about other clients in this party
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
	case "PartyChoose":
		var msg PartyChoose
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "PartyLeave":
		var msg PartyLeave
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	case "PartyCreate":
		var msg PartyCreate
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
func (pc PartyChoosing) MarshalJSON() ([]byte, error) {
	type alias PartyChoosing
	return fromTagMsg("PartyChoosing", alias(pc))
}

// MarshalJSON makes JSON.
func (pc PartyWaiting) MarshalJSON() ([]byte, error) {
	type alias PartyWaiting
	return fromTagMsg("PartyWaiting", alias(pc))
}
