package main

import (
	"encoding/json"
	"errors"
)

// ID is the type of IDs (client IDs, room IDs).
type ID uint64

// tagMsg is a JSON-encoded message. T denotes which type of thing to try to
// parse into, and P is the JSON encoding of that thing.
type tagMsg struct {
	T string
	P json.RawMessage
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

// PartyChoosing is the state of a client choosing their party.
type PartyChoosing struct {
	Name    string
	Parties []string
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

func (Close) isAction()      {}
func (NameChoose) isAction() {}

// Close means the client closed itself. No further Actions will follow from
// this client.
type Close struct{}

// NameChoose is a request from a client to choose their name.
type NameChoose struct {
	Name string
}

// ErrBadT means the T of a tagMsg did not match any known T.
var ErrBadT = errors.New("bad T")

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
	default:
		return nil, ErrBadT
	}
}

// IDAction is an Action from a particular client with a given ID.
type IDAction struct {
	ID
	Action
}
