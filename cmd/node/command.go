package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Doresimon/good-chain/console"
	HttpGoodRpc "github.com/Doresimon/good-chain/rpc/http"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/rpc/common"
	"gopkg.in/urfave/cli.v1"
)

// AppCommands ...
var appCommands = []cli.Command{
	{
		Name:        "start",
		Aliases:     []string{"run"},
		Usage:       "start a node",
		Description: "start a node",
		Action: func(c *cli.Context) error {
			// read config
			cfgBuffer, err := ioutil.ReadFile(configFile)
			if err != nil {
				console.Fatal(fmt.Sprintf("%s", err))
				return err
			}

			err = json.Unmarshal(cfgBuffer, &appConfig)
			if err != nil {
				console.Fatal(fmt.Sprintf("%s", err))
				return err
			}

			// read storage

			// construct state tree

			//

			path := appConfig.Chain

			C := new(chain.Chain)
			C.Genesis(path)
			C.I = 0

			ChainService := common.NewChainService()
			ChainService.I = 0
			ChainService.C = C

			ChainService.B = chain.NewBlock(C.BN())

			C.RunTicker(ChainService.B)

			console.Info("HttpGoodRpc.Server()")
			go HttpGoodRpc.Server(appConfig.Host, appConfig.Port, ChainService)

			ch := make(chan int) // block process
			<-ch
			return nil
		},
	},
}
