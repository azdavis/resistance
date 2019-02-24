package main

import (
	"encoding/json"
	"errors"
)

// ID is the type of IDs (client IDs, room IDs).
type ID uint64

// TagMsg is a JSON-encoded message. T denotes which type of thing to try to
// parse into, and P is the JSON encoding of that thing.
type TagMsg struct {
	T string
	P json.RawMessage
}

// FromTagMsg creates a JSON-encoded TagMsg.
func FromTagMsg(t string, p interface{}) ([]byte, error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return json.Marshal(TagMsg{T: t, P: json.RawMessage(bs)})
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
	return FromTagMsg("PartyChoosing", alias(pc))
}

// Action is a parsed TagMsg.P sent from a client. It is a request from the
// client to change state.
type Action interface {
	isAction()
}

func (NameChoose) isAction() {}

// NameChoose is a request from a client to choose their name.
type NameChoose struct {
	ID
	Name string
}

// ErrBadT means the T of a TagMsg did not match any known T.
var ErrBadT = errors.New("bad T")

// JSONToAction tries to turn a JSON encoding of a TagMsg into a Action.
func JSONToAction(bs []byte) (Action, error) {
	var tm TagMsg
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
