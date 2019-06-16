package main

import chain "github.com/Doresimon/good-chain/chain"

func main() {
	C := new(chain.Chain)
	C.Genesis("./chain.config")
}
