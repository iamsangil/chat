package models

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	TypeConnect = iota
	TypeSend
	TypeDisconnect
)

type MessageType int

type Message struct {
	Type     MessageType `json:"type,omitempty"`
	RoomID   string      `json:"room_id,omitempty"`
	ClientID string      `json:"client_id,omitempty"`
	Data     string      `json:"data,omitempty"`
}

func (m MessageType) MarshalJSON() ([]byte, error) {
	var msg string
	switch m {
	case TypeConnect:
		msg = "connect"
	case TypeSend:
		msg = "send"
	case TypeDisconnect:
		msg = "disconnect"
	default:
		msg = "invalid type"
	}
	return json.Marshal(msg)
}

func (m *MessageType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "\"connect\"":
		*m = TypeConnect
	case "\"send\"":
		*m = TypeSend
	case "\"disconnect\"":
		*m = TypeDisconnect
	default:
		return errors.New("invalid type")
	}
	return nil
}

func (m MessageType) String() string {
	var str string
	switch m {
	case TypeConnect:
		str = "connect"
	case TypeSend:
		str = "send"
	case TypeDisconnect:
		str = "disconnect"
	default:
		str = "invalid type"
	}
	return str
}
