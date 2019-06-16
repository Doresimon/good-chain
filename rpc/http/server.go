package rpc

import (
	"net"
	"net/http"
	"net/rpc"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
	ER "github.com/Doresimon/good-chain/error"
	"github.com/Doresimon/good-chain/rpc/common"
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

	C.RunTicker(ChainService.B)

	rpc.Register(ChainService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", port)
	ER.Check("net.Listen() failed", e)

	// go http.Serve(l, nil)
	http.Serve(l, nil)

	console.Dev("terminated!")
}
