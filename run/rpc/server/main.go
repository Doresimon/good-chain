package main

import (
	"good-chain/console"
	HttpGoodRpc "good-chain/rpc/http"
)

func main() {
	console.Info("HttpGoodRpc.Server()")

	HttpGoodRpc.Server()
}
