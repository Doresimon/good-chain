package main

import (
	"good-chain/console"
	"good-chain/rpc/common"

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
	args.Data = []string{"Alice", "0", "Hello world"}
	// args.Data = []string{"1", "2"}
	var result string

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
