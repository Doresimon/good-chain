package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"

	"github.com/Doresimon/good-chain/console"
)

var app = cli.NewApp()

func main() {
	console.ShowColors()

	app := App()

	console.Info("start app")
	err := app.Run(os.Args)

	if err != nil {
		console.Fatal("error")
	}

}
