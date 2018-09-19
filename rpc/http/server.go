package rpc

import (
	"fmt"
	"good-chain/rpc/common"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Server ...
// run a http server,
func Server() {
	port := ":1234"
	ChainService := common.NewChainService()
	rpc.Register(ChainService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", port)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	// go http.Serve(l, nil)
	http.Serve(l, nil)

	fmt.Println("terminated!")
}
