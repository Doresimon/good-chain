package main

import (
	"os"

	"github.com/Doresimon/good-chain/console"
)

func main() {
	setupKey()
	startP2P()

	app := App()
	err := app.Run(os.Args)
	if err != nil {
		console.Fatal("error")
	}
}
