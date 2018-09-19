package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	C "good-chain/crypto"
)

// Log ...
type Log struct {
	Sender    string       `json:"sender"`
	SupposeBN string       `json:"suppose Block number"`
	Message   string       `json:"message"`
	Sig       *C.Signature `json:"signature"`
	Hash      string       `json:"hash"`
}

// NewLog ...
func NewLog(s string, sBN string, m string, sig string) *Log {
	// check sig
	// pk, _ := hex.DecodeString(s)

	L := new(Log)
	L.Sender = s
	L.SupposeBN = sBN
	L.Message = m

	L.Sig = new(C.Signature)
	json.Unmarshal([]byte(sig), L.Sig)

	_sig, _ := json.Marshal(L.Sig)

	sum := sha256.Sum256([]byte(L.Sender + L.SupposeBN + L.Message + string(_sig)))
	L.Hash = hex.EncodeToString(sum[:])

	return L
}
