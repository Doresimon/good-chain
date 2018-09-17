package rpc

import (
	"fmt"
	"net"
	"net/rpc"

	"../common"
)

func Server() {
	var ms = new(common.MathService)
	rpc.Register(ms)
	fmt.Println("start service...")
	var address, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1234")
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		fmt.Println("start failed:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("receive a request...")
		rpc.ServeConn(conn)
	}
	fmt.Println("terminated!")
}
