package rpc

import (
	"fmt"
	"net/rpc"
)

// HTTPClient ...
type HTTPClient struct {
	client  *rpc.Client
	network string
	address string
}

// Call ...
// call a method from server
func (c *HTTPClient) Call(m string, args interface{}) (string, error) {
	fmt.Println("rpc.HTTPClient.Call()")
	result := ""

	err := c.client.Call(m, args, &result)
	if err != nil {
		fmt.Println("call failed: ", err)
	}
	return result, err
}

// NewClient ...
// create a new Http Client
func NewClient(n string, a string) (*HTTPClient, error) {
	fmt.Println("rpc.NewClient()")

	var err error

	c := new(HTTPClient)
	c.network = n
	c.address = a
	c.client, err = rpc.DialHTTP(c.network, c.address)

	if err != nil {
		fmt.Println("dail failed: ", err)
	}

	return c, err
}
