package main

import (
	"fmt"

	HttpGoodRpc "../../../rpc/http"
	// TcpGoodRpc "../../../rpc/tcp"
)

func main() {
	fmt.Printf("HttpGoodRpc.Client()\n")

	HttpGoodRpc.Client()
}