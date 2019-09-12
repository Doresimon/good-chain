package main

import (
	"time"

	"github.com/Doresimon/good-chain/p2p"
	"github.com/urfave/cli"
)

// AppCommands ...
var appCommands = []cli.Command{
	{
		Name:        "new-org",
		Aliases:     []string{"new-org"},
		Usage:       "create a new org account",
		Description: "create a new org account",
		Action: func(c *cli.Context) error { // read config
			go func() {
				logBytes := newOrg()

				msg := p2p.NewMessage(p2p.LOG, logBytes)
				data := p2p.Serialize(msg)
				_, err := rw.Write(data)
				if err != nil {
					panic(err)
				}
				err = rw.Flush()
				if err != nil {
					panic(err)
				}
			}()

			time.Sleep(time.Second * 20)

			go func() {
				logBytes := newAccount()

				msg := p2p.NewMessage(p2p.LOG, logBytes)
				data := p2p.Serialize(msg)
				_, err := rw.Write(data)
				if err != nil {
					panic(err)
				}
				err = rw.Flush()
				if err != nil {
					panic(err)
				}
			}()

			select {}
		},
	},
}

// App ...
func App() *cli.App {
	var app = cli.NewApp()

	app.Name = "node"
	app.Usage = "run a node of good chain"
	app.Commands = appCommands

	return app
}
