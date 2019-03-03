package main

import (
	"encoding/json"
	"errors"
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

// ToClient //////////////////////////////////////////////////////////////////

// ToClient is a state that a client could be in. It is sent to the client to
// change the client's state. It may be sent in direct reply to a client's
// ToServer, or it may be sent because the client was transitively affected by
// another client's ToServer.
type ToClient interface {
	json.Marshaler
	isToClient()
}

func (PartyChoosing) isToClient() {}
func (PartyWaiting) isToClient()  {}

// PartyChoosing is the state of a client choosing their party.
type PartyChoosing struct {
	Parties []PartyInfo // available parties to join
}

// PartyWaiting is the state of a client who is in a party, but the game has not
// yet started.
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

// ErrUnknownActionType means the T of a tagMsg did not match any known T.
var ErrUnknownActionType = errors.New("unknown action type")

// JSONToAction tries to turn a JSON encoding of a tagMsg into a ToServer.
func JSONToAction(bs []byte) (ToServer, error) {
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
	default:
		return nil, ErrUnknownActionType
	}
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
