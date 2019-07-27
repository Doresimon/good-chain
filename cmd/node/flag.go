package main

import (
	"gopkg.in/urfave/cli.v1"
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
