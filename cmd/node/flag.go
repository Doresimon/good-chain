package main

import (
	"github.com/urfave/cli"
)


// appFlags ...
var appFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "config, c",     //
		Value:       "./config.json", // default
		Usage:       "config file path",
		Destination: &configFile,
	},
}
