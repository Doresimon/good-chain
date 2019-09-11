package main

import (
	"os"

	"github.com/Doresimon/good-chain/console"
)

// var app = cli.NewApp()

func main() {
	console.ShowColors()

	app := App()
	err := app.Run(os.Args)

	if err != nil {
		console.Fatal("error")
	}
}
