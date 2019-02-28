package main

import (
	"encoding/json"
	"errors"
)

// tagMsg is a JSON-encoded message. T denotes which type of thing to try to
// parse into, and P is the JSON encoding of that thing.
type tagMsg struct {
	T string          // type of thing
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

// State is a state that a client could be in. It is sent to the client to
// change the client's state. It may be sent in direct reply to a client's
// Action, or it may be sent because the client was transitively affected by
// another client's Action.
type State interface {
	json.Marshaler
	isState()
}

func (PartyChoosing) isState() {}

// PartyInfo contains a party name (may not be unique) and PID (unique).
type PartyInfo struct {
	PID
	Name string
}

// PartyChoosing is the state of a client choosing their party.
type PartyChoosing struct {
	Name    string      // name of client
	Parties []PartyInfo // available parties to join
}

// MarshalJSON makes JSON.
func (pc PartyChoosing) MarshalJSON() ([]byte, error) {
	type alias PartyChoosing
	return fromTagMsg("PartyChoosing", alias(pc))
}

// Action is a usually parsed tagMsg.P sent from a client. It is a request from
// the client to change state. The only Action that does not originate from a
// tagMsg is Close. The client "sends" a Close Action by closing itself.
type Action interface {
	isAction()
}

func (Close) isAction()       {}
func (NameChoose) isAction()  {}
func (PartyChoose) isAction() {}
func (PartyCreate) isAction() {}

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

// PartyCreate is a request to create a new party, with oneself as the leader.
type PartyCreate struct {
	Name string // desired party name
}

// ErrUnknownActionType means the T of a tagMsg did not match any known T.
var ErrUnknownActionType = errors.New("unknown action type")

// JSONToAction tries to turn a JSON encoding of a tagMsg into a Action.
func JSONToAction(bs []byte) (Action, error) {
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
	case "PartyCreate":
		var msg PartyCreate
		err = json.Unmarshal(tm.P, &msg)
		return msg, err
	default:
		return nil, ErrUnknownActionType
	}
}
