package rpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/Doresimon/good-chain/console"
	ERROR "github.com/Doresimon/good-chain/error"
	"github.com/Doresimon/good-chain/rpc/common"
)

// Server ...
// run a http server,
func Server(address string, port uint, chainService *common.ChainService) {

	rpc.Register(chainService)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	ERROR.Check("net.Listen() failed", err)

	// go http.Serve(l, nil)
	http.Serve(l, nil)

	console.Dev("terminated!")
}
