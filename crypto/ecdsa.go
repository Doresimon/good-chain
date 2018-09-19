package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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
