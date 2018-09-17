package rpc

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"../common"
)

// Server ...
// run a http server,
func Server() {
	port := ":1234"
	ChainService := new(common.ChainService)
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
