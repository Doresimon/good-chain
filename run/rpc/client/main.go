package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"good-chain/console"
	C "good-chain/crypto"
	E "good-chain/error"
	"good-chain/rpc/common"
	"strconv"
	"strings"

	HttpGoodRpc "good-chain/rpc/http"
)

func main() {
	CL()
	// CL()
	// CL()
	// CL()
	// CL()
	// CL()
	// CL()
	GP()
}

func CL() {
	var args = new(common.Args)
	args.Data = []string{"[pk]", "[block number]", "message", "[sig]"}
	// args.Data = []string{"1", "2"}
	var result string

	// sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// sk.PublicKey

	sk, err := C.GenerateKey()
	E.Check("ecdsa.GenerateKey", err)
	pk := sk.S.PublicKey
	pkHex := hex.EncodeToString(C.Marshal(pk))

	h := sha256.Sum256([]byte(strings.Join(args.Data[0:2], ",")))
	r, s, err := ecdsa.Sign(rand.Reader, sk.S, h[:])
	sig := C.NewSignature(r, s, h[:])
	sigstr, _ := sig.Marshal()

	args.Data[0] = pkHex
	args.Data[1] = strconv.FormatInt(1, 10)
	args.Data[2] = "hello WORLD"
	args.Data[3] = string(sigstr)

	var method = "ChainService.CreateLog"

	c, err := HttpGoodRpc.NewClient("tcp", "127.0.0.1:1234")

	if err != nil {
		console.Error("HttpGoodRpc.NewClient()")
		return
	}

	result, _ = c.Call(method, args)

	console.Info("result:" + result)
}

func GP() {
	var args = new(common.Args)
	args.Data = []string{""}
	// args.Data = []string{"1", "2"}
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
