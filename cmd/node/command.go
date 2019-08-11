package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/p2p"

	"github.com/Doresimon/good-chain/chain"
	"github.com/urfave/cli"
)

// AppCommands ...
var appCommands = []cli.Command{
	{
		Name:        "start",
		Aliases:     []string{"run", "chain-node"},
		Usage:       "start a node",
		Description: "start a node",
		Action: func(c *cli.Context) error { // read config
			cfgBuffer, err := ioutil.ReadFile(configFile)
			if err != nil {
				console.Fatal(err.Error())
				return err
			}

			err = json.Unmarshal(cfgBuffer, &appConfig)
			if err != nil {
				console.Fatal(err.Error())
				return err
			}

			var nodeService struct {
				cs *chain.Service
				ps *p2p.Service
			}

			path := appConfig.Chain
			chainInstance := chain.NewChain(path)
			chainService := chain.NewService(chainInstance)
			p2pService := p2p.NewService(chainService)

			nodeService.cs = chainService
			nodeService.ps = p2pService

			select {}
		},
	},
}
