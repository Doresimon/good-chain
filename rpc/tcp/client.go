package rpc

import (
	"fmt"
	"net/rpc"

	"../common"
)

func Client() {
	var args = common.Args{40, 3}
	var result = common.Result{}

	var client, err = rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("failed to connect RPC server via TCP:", err)
	}
	fmt.Println("[CALL] start calling")
	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Println("call failed:", err)
	}
	fmt.Println("call result:", result.Value)
}
