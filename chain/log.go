package chain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
	"reflect"

	C "github.com/Doresimon/good-chain/crypto"
)

// Log ...
type Log struct {
	Sender    []byte       `json:"sender"`
	SupposeBN []byte       `json:"suppose Block number"`
	Message   []byte       `json:"message"`
	Sig       *C.Signature `json:"signature"`
	Hash      []byte       `json:"hash"`
}

// NewLog ...
func NewLog(Sender []byte, SupposeBN []byte, Message []byte, R []byte, S []byte, H []byte) *Log {
	// check sig
	// pk, _ := hex.DecodeString(s)

	L := new(Log)
	L.Sender = Sender
	L.SupposeBN = SupposeBN
	L.Message = Message

	L.Sig = new(C.Signature)

	L.Sig.R = *new(big.Int)
	L.Sig.S = *new(big.Int)
	L.Sig.H = H

	L.Sig.R.SetBytes(R)
	L.Sig.S.SetBytes(S)

	// sum := sha256.Sum256(L.Sender + L.SupposeBN + L.Message + L.Sig.R + L.Sig.S)
	sum := sha256.Sum256(append(append(append(append(L.Sender, L.SupposeBN...), L.Message...), R...), S...))
	L.Hash = sum[:]

	return L
}

func (L *Log) VerifySig() bool {
	pk := C.UnMarshalPK(L.Sender)
	hash := sha256.Sum256(append(append(L.Sender, L.SupposeBN...), L.Message...))

	if !reflect.DeepEqual(L.Sig.H, hash[:]) {
		return false
	}
	return ecdsa.Verify(pk, hash[:], &L.Sig.R, &L.Sig.S)
}
