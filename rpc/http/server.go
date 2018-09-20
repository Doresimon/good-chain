package rpc

import (
	"good-chain/chain"
	"good-chain/console"
	ER "good-chain/error"
	"good-chain/rpc/common"
	"net"
	"net/http"
	"net/rpc"
)

// Server ...
// run a http server,
func Server() {
	port := ":1234"

	path := "./chain.config"

	C := new(chain.Chain)
	C.Genesis(path)
	C.I = 0

	ChainService := common.NewChainService()
	ChainService.I = 0
	ChainService.C = C

	ChainService.B = chain.NewBlock(C.BN())

	C.RunTicker()

	rpc.Register(ChainService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", port)
	ER.Check("net.Listen() failed", e)

	// go http.Serve(l, nil)
	http.Serve(l, nil)

	console.Dev("terminated!")
}
