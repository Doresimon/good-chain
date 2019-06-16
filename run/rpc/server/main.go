package main

import (
	"github.com/Doresimon/good-chain/console"
	HttpGoodRpc "github.com/Doresimon/good-chain/rpc/http"
)

func main() {
	console.Info("HttpGoodRpc.Server()")

	HttpGoodRpc.Server()
}
