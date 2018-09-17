package main

import "good-chain/chain"

func main() {
	C := new(chain.Chain)
	C.Genesis("./chain.config")
}
