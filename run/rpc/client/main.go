package main

import (
	"fmt"
	"good-chain/rpc/common"

	HttpGoodRpc "good-chain/rpc/http"
)

func main() {
	L()
}

func L() {
	var args = new(common.Args)
	// args.Data = []string{"1", "2", "3", "4", "5"}
	args.Data = []string{"1", "2"}
	var result string

	var method = "ChainService.CreateLog"

	c, _ := HttpGoodRpc.NewClient("tcp", "127.0.0.1:1234")

	result, _ = c.Call(method, args)

	fmt.Printf("result: %v\n", result)
}
