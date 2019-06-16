package main

import "github.com/Doresimon/good-chain"

func main() {
	C := new(chain.Chain)
	C.Genesis("./chain.config")
}
