package main

import (
	"gopkg.in/urfave/cli.v1"
)

type config struct {
	Port  uint   `json:"port"`
	Host  string `json:"host"`
	Chain string `json:"chain"`
}

var configFile string
var appConfig config

// App ...
func App() *cli.App {
	var app = cli.NewApp()

	app.Name = "node"
	app.Usage = "run a node of good chain"
	app.Flags = appFlags
	app.Commands = appCommands

	return app
}
