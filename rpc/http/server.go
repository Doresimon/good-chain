package rpc

import (
	"fmt"
	"net/http"
	"net/rpc"

	"../common"
)

func Server() {
	port := ":1234"
	var ms = new(common.MathService)
	rpc.Register(ms)
	rpc.HandleHTTP() //将Rpc绑定到HTTP协议上。
	fmt.Println("listening port "+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("服务已停止!")
}