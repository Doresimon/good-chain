package hdk

import (
	"fmt"
	"math/big"

	"github.com/Doresimon/good-chain/crypto/bls"
	"github.com/Doresimon/good-chain/crypto/hash/hmac"
	"golang.org/x/crypto/bn256"
)

var bn256Order = bn256.Order
var bigZero = new(big.Int).SetInt64(0)
var bigTmp = new(big.Int).SetInt64(0)

// // HDPrivateKey ...
// type HDPrivateKey interface {
// 	Public() *HDPublicKey
// 	Sign([]byte) ([]byte, error)
// 	Bytes() []byte
// }

// // HDPublicKey ...
// type HDPublicKey interface {
// 	Bytes() []byte
// 	Verify(message, sigBytes []byte) bool
// }

// GenerateMasterKey TODO
func GenerateMasterKey(key, seed []byte) (masterKey *bls.PrivateKey, chainCode []byte, err error) {
	MAC := hmac.SHA512(key, seed)
	privateKeyValue := new(big.Int).SetBytes(MAC[0:32])
	chainCode = MAC[32:64]

	privateKeyValue.Mod(privateKeyValue, bn256Order)

	privateKeyValueBytes := privateKeyValue.Bytes()
	len := len(privateKeyValueBytes)
	if len < 32 {
		emptyBytes := make([]byte, 32-len)
		privateKeyValueBytes = append(emptyBytes, privateKeyValueBytes...)
	}
	if len > 32 {
		panic("len > 32") // it should never go here
	}

	privateKeyValue.SetBytes(privateKeyValueBytes)
	if privateKeyValue.Sign() == 0 {
		err = fmt.Errorf("master key is all 0")
	}

	masterKey = new(bls.PrivateKey)
	masterKey.Set(privateKeyValue)
	return
}

// Priv2Priv The function Priv2Priv((k_p, c_p), i) → (k_c, c_c) computes a child extended private key from the parent extended private key
func Priv2Priv(parentPrivKey *bls.PrivateKey, parentChainCode []byte, index uint32) (*bls.PrivateKey, []byte, bool) {
	key := parentChainCode
	data := make([]byte, 0, 0) // Data = ser(point(k)) || ser (i))

	parentPubKey := new(bn256.G2).ScalarBaseMult(parentPrivKey.Value())
	d1 := parentPubKey.Marshal()
	d2 := int32ToBytes(index)
	data = append(data, d1...)
	data = append(data, d2...)

	mac := hmac.SHA512(key, data)

	childKeyValue := new(big.Int)
	childKeyValue.SetBytes(mac[0:32])
	childKeyValue.Add(childKeyValue, parentPrivKey.Value())
	childKeyValue.Mod(childKeyValue, bn256Order)

	childChainCode := make([]byte, 32, 32)
	copy(childChainCode, mac[32:64])

	if childKeyValue.Sign() == 0 {
		return nil, nil, false
	}

	childKey := new(bls.PrivateKey)
	childKey.Set(childKeyValue)
	return childKey, childChainCode, true
}

// Pub2Pub The function Pub2Pub((K , c ), i) → (K , c ) computes a child extended public key from
// the parent extended public key
func Pub2Pub(parentPubKey *bls.PublicKey, parentChainCode []byte, index uint32) (*bls.PublicKey, []byte, bool) {
	key := parentChainCode
	data := make([]byte, 0, 0) // Data = []byte(point(k)) || [4]byte(i))

	d1 := parentPubKey.Value().Marshal()
	d2 := int32ToBytes(index)
	data = append(data, d1...)
	data = append(data, d2...)

	mac := hmac.SHA512(key, data)

	tmpBN := new(big.Int)
	tmpBN.SetBytes(mac[0:32])
	tmpBN.Mod(tmpBN, bn256Order)
	if tmpBN.Sign() == 0 {
		return nil, nil, false
	}
	// if bigTmp.Sub(bn256Order, tmpBN).Sign() != 1 {
	// 	return nil, nil, false
	// }

	childPubKeyValue := new(bn256.G2)
	childPubKeyValue.ScalarBaseMult(tmpBN)

	childPubKeyValue.Add(childPubKeyValue, parentPubKey.Value())

	childChainCode := make([]byte, 32, 32)
	copy(childChainCode, mac[32:64])

	childPubKey := new(bls.PublicKey)
	childPubKey.Set(childPubKeyValue)

	return childPubKey, childChainCode, true
}

func int32ToBytes(i uint32) []byte {
	ret := make([]byte, 4, 4)
	// ret[0] = byte((i & 0xff000000) >> 24)
	// ret[1] = byte((i & 0x00ff0000) >> 16)
	// ret[2] = byte((i & 0x0000ff00) >> 8)
	// ret[3] = byte((i & 0x000000ff) >> 0)
	ret[0] = byte((i) >> 24)
	ret[1] = byte((i) >> 16)
	ret[2] = byte((i) >> 8)
	ret[3] = byte((i) >> 0)
	return ret
}

func isBytesEmpty(b *[]byte) bool {
	empty := true
	for _, v := range *b {
		if v != 0 {
			empty = false
			break
		}
	}
	return empty
}
