package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/Doresimon/good-chain/crypto/bls"
)

var masterPrivKey *bls.PrivateKey
var masterChainCode []byte

func setupKey() {
	// key setup
	// key := []byte("good chain key")
	// seed, _ := rand.Bytes(256)
	// var err error
	// masterPrivKey, masterChainCode, err = hdk.GenerateMasterKey(key, seed)
	// if err != nil {
	// 	panic(err)
	// }
	privHex := "63d252abaa4d98ecd978a0eb81c4624d8beebfb4c2e3817786da6ed6685cca04"
	privBytes, err := hex.DecodeString(privHex)
	if err != nil {
		panic(err)
	}
	masterPrivKey = new(bls.PrivateKey).Set(new(big.Int).SetBytes(privBytes))

	masterChainCodeHex := "1c6ee42d6341af5ee099669379220e9aa01973f70c5ba8df99a31234433b90e6"
	masterChainCode, err := hex.DecodeString(masterChainCodeHex)
	if err != nil {
		panic(err)
	}

	fmt.Printf("masterPrivKey   = %x\n", masterPrivKey.Bytes())
	fmt.Printf("masterChainCode = %x\n", masterChainCode)
}
