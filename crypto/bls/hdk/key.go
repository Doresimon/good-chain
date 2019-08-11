package hdk

import (
	"fmt"
	"math/big"

	"github.com/Doresimon/good-chain/crypto/hash/hmac"
	"golang.org/x/crypto/bn256"
)

var bn256Order = bn256.Order
var bigZero = new(big.Int).SetInt64(0)
var bigTmp = new(big.Int).SetInt64(0)

// GenerateMasterKey TODO
func GenerateMasterKey(key, seed []byte) (masterKey *big.Int, chainCode []byte, err error) {
	MAC := hmac.SHA512(key, seed)
	masterKey = new(big.Int).SetBytes(MAC[0:32])
	chainCode = MAC[32:64]

	masterKey.Mod(masterKey, bn256Order)

	masterKeyBytes := masterKey.Bytes()
	len := len(masterKeyBytes)
	if len < 32 {
		emptyBytes := make([]byte, 32-len)
		masterKeyBytes = append(emptyBytes, masterKeyBytes...)
	}
	if len > 32 {
		panic("len > 32") // it should never go here
	}

	masterKey.SetBytes(masterKeyBytes)
	if masterKey.Sign() == 0 {
		err = fmt.Errorf("master key is all 0")
	}
	return
}

// Priv2Priv The function Priv2Priv((k_p, c_p), i) → (k_c, c_c) computes a child extended private key from the parent extended private key
func Priv2Priv(parentPrivKey *big.Int, parentChainCode []byte, index uint32) (*big.Int, []byte, bool) {
	key := parentChainCode
	data := make([]byte, 0, 0) // Data = ser(point(k)) || ser (i))

	parentPubKey := new(bn256.G2).ScalarBaseMult(parentPrivKey)
	d1 := parentPubKey.Marshal()
	d2 := int32ToBytes(index)
	data = append(data, d1...)
	data = append(data, d2...)

	mac := hmac.SHA512(key, data)

	childKey := new(big.Int)
	childKey.SetBytes(mac[0:32])
	childKey.Add(childKey, parentPrivKey)
	childKey.Mod(childKey, bn256Order)

	childChainCode := make([]byte, 32, 32)
	copy(childChainCode, mac[32:64])

	if childKey.Sign() == 0 {
		return nil, nil, false
	}
	// if bigTmp.Sub(bn256Order, childKey).Sign() != 1 {
	// 	return nil, nil, false
	// }
	return childKey, childChainCode, true
}

// Pub2Pub The function Pub2Pub((K , c ), i) → (K , c ) computes a child extended public key from
// the parent extended public key
func Pub2Pub(parentPubKey *bn256.G2, parentChainCode []byte, index uint32) (*bn256.G2, []byte, bool) {
	key := parentChainCode
	data := make([]byte, 0, 0) // Data = []byte(point(k)) || [4]byte(i))

	d1 := parentPubKey.Marshal()
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

	childPubKey := new(bn256.G2)
	childPubKey.ScalarBaseMult(tmpBN)

	childPubKey.Add(childPubKey, parentPubKey)

	childChainCode := make([]byte, 32, 32)
	copy(childChainCode, mac[32:64])

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
