package rpc

import (
	"net/rpc"

	"github.com/Doresimon/good-chain/console"
	ER "github.com/Doresimon/good-chain/error"
)

// HTTPClient ...
type HTTPClient struct {
	client  *rpc.Client
	network string
	address string
}

// Call ...
// call a method from server
func (c *HTTPClient) Call(method string, args interface{}) (string, error) {
	console.Dev("rpc.HTTPClient.Call()")
	result := ""

	err := c.client.Call(method, args, &result)
	ER.Check("call failed", err)

	return result, err
}

// NewClient ...
// create a new Http Client
func NewClient(n string, a string) (*HTTPClient, error) {
	console.Dev("rpc.NewClient()")

	var err error

	c := new(HTTPClient)
	c.network = n
	c.address = a
	c.client, err = rpc.DialHTTP(c.network, c.address)
	ER.Check("dial failed", err)

	return c, err
}
