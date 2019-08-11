package chain

import (
	"crypto/sha256"
	"encoding/json"
)

// NewLog ...
func NewLog(sender []byte, bn []byte, tx []byte, sig []byte) *Log {
	hash := sha256.Sum256(tx)

	L := new(Log)
	L.Sender = sender
	L.SupposeBN = bn
	L.Sig = sig
	L.TX = tx
	L.Hash = hash[:]

	return L
}

// Log ...
type Log struct {
	Sender    []byte `json:"sender"`
	SupposeBN []byte `json:"suppose-block-number"`
	TX        []byte `json:"transaction"`
	Sig       []byte `json:"signature"`
	Hash      []byte `json:"hash"`
	// Sig       *Gcrypto.Signature `json:"signature"`
	// Message   []byte             `json:"message"`
}

// Marshal is a wrapper of type Log
func (l *Log) Marshal() ([]byte, error) {
	return json.Marshal(l)
}

// UnmarshalLog parse a byte slice of a log
func UnmarshalLog(buf []byte) (*Log, error) {
	l := new(Log)
	err := json.Unmarshal(buf, l)
	return l, err
}

// type LogPool chan *Log

// // NewLog ...
// func NewLog(Sender []byte, SupposeBN []byte, Message []byte, R []byte, S []byte, H []byte) *Log {
// 	// check sig
// 	// pk, _ := hex.DecodeString(s)

// 	L := new(Log)
// 	L.Sender = Sender
// 	L.SupposeBN = SupposeBN
// 	L.Message = Message

// 	L.Sig = new(Gcrypto.Signature)

// 	L.Sig.R = *new(big.Int)
// 	L.Sig.S = *new(big.Int)
// 	L.Sig.H = H

// 	L.Sig.R.SetBytes(R)
// 	L.Sig.S.SetBytes(S)

// 	// sum := sha256.Sum256(L.Sender + L.SupposeBN + L.Message + L.Sig.R + L.Sig.S)
// 	sum := sha256.Sum256(append(append(append(append(L.Sender, L.SupposeBN...), L.Message...), R...), S...))
// 	L.Hash = sum[:]

// 	return L
// }

// VerifySig ...
// func (L *Log) VerifySig() bool {
// 	pk := Gcrypto.UnMarshalPK(L.Sender)
// 	hash := sha256.Sum256(append(append(L.Sender, L.SupposeBN...), L.Message...))

// 	if !reflect.DeepEqual(L.Sig.H, hash[:]) {
// 		return false
// 	}
// 	return ecdsa.Verify(pk, hash[:], &L.Sig.R, &L.Sig.S)
// }
