package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/Doresimon/good-chain/crypto/coding"
)

const (
	HEARTBEAT = iota
	SYNC
	HELLO
	LOG
)

// Message is one message transfered throught p2p channel
type Message struct {
	Type    int    `json:"type"`
	Content []byte `json:"content"`
}

// NewMessage create a new message instance
func NewMessage(t int, c []byte) *Message {
	m := new(Message)
	m.Type = t
	m.Content = c
	return m
}

// Serialize makes message as byte slice with length as prefix
func (m *Message) Serialize() *[]byte {
	buf := Serialize(m)
	return &buf
}

// Serialize makes message as byte slice with length as prefix
func Serialize(m *Message) []byte {
	mBytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	buf := coding.Uint32ToBytes(uint32(len(mBytes)))
	buf = append(buf, mBytes...)

	return buf
}

// Unserialize un serialize a byte slice to a message
func Unserialize(buf []byte) (*Message, error) {
	m := new(Message)

	bugLen := len(buf)
	if bugLen < 4 {
		return nil, fmt.Errorf("buf is broken")
	}

	msgLen := coding.BytesToUint32(buf[0:4])
	if bugLen < int(msgLen+4) {
		panic("raw message is broken")
	}

	err := json.Unmarshal(buf[4:4+msgLen], m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
