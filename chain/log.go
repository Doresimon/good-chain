package chain

import (
	"crypto/sha256"
	"encoding/json"
)

// NewLog ...
func NewLog(sender []byte, sig []byte, body *Body) *Log {
	L := new(Log)
	L.Sender = sender
	L.Sig = sig
	L.Body = body

	return L
}

// Log ...
type Log struct {
	Sender []byte `json:"sender"`    // 发送者公钥
	Sig    []byte `json:"signature"` // 发送者对消息体签名
	Body   *Body  `json:"body"`      // 消息体
}

// Marshal is a wrapper of json.Marshal for type Log
func (l *Log) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

// Hash returns the hash of body
func (l *Log) Hash() []byte {
	logBodyBytes, _ := json.Marshal(l.Body)
	hash := sha256.Sum256(logBodyBytes)
	return hash[:]
}

// UnmarshalLog parse a byte slice of a log
func UnmarshalLog(buf []byte) (*Log, error) {
	l := new(Log)
	err := json.Unmarshal(buf, l)
	return l, err
}
