package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"math/big"
)

type Secret struct {
	S *ecdsa.PrivateKey
}

func GenerateKey() (*Secret, error) {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	SK := new(Secret)
	SK.S = sk
	return SK, err
}

func Marshal(pk ecdsa.PublicKey) []byte {
	return elliptic.Marshal(pk.Curve, pk.X, pk.Y)
}

// func UnMarshal(pk []byte) []byte {
// 	return elliptic.Unmarshal(pk)
// }

type Signature struct {
	H []byte  `json:"hash"`
	R big.Int `json:"r"`
	S big.Int `json:"s"`
}

func NewSignature(r, s *big.Int, h []byte) *Signature {
	sig := new(Signature)
	sig.R = *r
	sig.S = *s
	sig.H = h

	return sig
}

func (this *Signature) Marshal() ([]byte, error) {
	return json.Marshal(this)
}
