package main

import (
	"crypto"
	"crypto/rand"
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
	CL()
	CL()
	CL()
	CL()
	CL()
	CL()
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
	sig, err := sk.S.Sign(rand.Reader, []byte(strings.Join(args.Data, ",")), crypto.SHA256)
	sigHex := hex.EncodeToString(sig)

	args.Data[0] = pkHex
	args.Data[1] = strconv.FormatInt(1, 10)
	args.Data[2] = "hello WORLD"
	args.Data[3] = sigHex

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
