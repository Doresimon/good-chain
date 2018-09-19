package main

import (
	"fmt"

	HttpGoodRpc "good-chain/rpc/http"
)

func main() {
	fmt.Printf("HttpGoodRpc.Server()\n")

	HttpGoodRpc.Server()
}
