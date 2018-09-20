package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"good-chain/console"
	C "good-chain/crypto"
	E "good-chain/error"
	"good-chain/rpc/common"
	"os"
	"strconv"

	HttpGoodRpc "good-chain/rpc/http"
)

func main() {

	argsWithProg := os.Args

	fmt.Println(argsWithProg)

	switch argsWithProg[1] {
	case "new":
		NewMessage()
		break
	case "getpool":
		GetPool()
		break
	case "getblock":
		GetBlock(argsWithProg[2])
		break
	default:
		GetPool()
	}
}

func NewMessage() {
	var args = new(common.HexMessage)

	var result string

	sk, err := C.GenerateKey()
	E.Check("ecdsa.GenerateKey", err)
	pk := sk.S.PublicKey

	bn := int64(1)
	message := "HELLO WORLD"

	pk_bytes := []byte(C.MarshalPK(pk))
	bn_bytes := []byte(strconv.FormatInt(bn, 10))
	message_bytes := []byte(message)

	args.Pk = hex.EncodeToString(pk_bytes)
	args.SupposeBlockNumber = hex.EncodeToString(bn_bytes)
	args.Message = hex.EncodeToString(message_bytes)

	h := sha256.Sum256(append(append(pk_bytes, bn_bytes...), message_bytes...))
	r, s, err := ecdsa.Sign(rand.Reader, sk.S, h[:])
	// sig := C.NewSignature(r, s, h[:])
	// r.SetUint64(100)

	args.Sig = *new(common.StringSig)
	args.Sig.R = r.Bytes()
	args.Sig.S = s.Bytes()
	args.Sig.H = h[:]

	var method = "ChainService.NewLog"

	c, err := HttpGoodRpc.NewClient("tcp", "127.0.0.1:1234")

	if err != nil {
		console.Error("HttpGoodRpc.NewClient()")
		return
	}

	// for i := 0; i < 10000; i++ {

	// 	result, _ = c.Call(method, args)
	// }
	result, _ = c.Call(method, args)

	console.Info("result:" + result)
}

func GetPool() {
	var args = new(common.Args)
	args.Data = []string{""}
	var result string

	var method = "ChainService.GetPool"

	c, err := HttpGoodRpc.NewClient("tcp", "127.0.0.1:1234")

	if err != nil {
		console.Error("HttpGoodRpc.NewClient()")
		return
	}

	result, _ = c.Call(method, args)

	console.Info("result:" + result)
}

func GetBlock(arg string) {
	BN, err := strconv.Atoi(arg)

	E.Check("argument is wrong", err)

	var result string

	var method = "ChainService.GetBlock"

	c, err := HttpGoodRpc.NewClient("tcp", "127.0.0.1:1234")

	if err != nil {
		console.Error("HttpGoodRpc.NewClient()")
		return
	}

	result, _ = c.Call(method, uint64(BN))

	console.Info("result:" + result)
}
