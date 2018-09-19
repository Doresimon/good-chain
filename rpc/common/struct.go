package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"good-chain/chain"
	"good-chain/console"
)

// Args ...
// ajax data
type Args struct {
	Data []string
}

// ChainService ...
type ChainService struct {
	c *chain.Chain
	b *chain.Block
}

// CreateBlock ...
func (cs *ChainService) CreateBlock(args *Args, result *string) error {
	console.Dev("ChainService.CreateBlock()")

	for s := range args.Data {
		fmt.Println(s)
	}

	*result = "hiahiahia"
	return nil
}

// CreateLog ...
func (cs *ChainService) CreateLog(args *Args, result *string) error {
	console.Dev("ChainService.CreateLog()")

	// for s := range args.Data {
	// 	fmt.Println(s)
	// }
	if len(args.Data) < 3 {
		*result = "[ERROR] not enough arguments"
		return errors.New("[ERROR] not enough arguments")
	}

	s := args.Data[0]
	sBN := args.Data[1]
	m := args.Data[2]

	L := chain.NewLog(s, sBN, m)

	cs.b.AddLog(L)

	Ls, _ := json.Marshal(L)

	*result = string(Ls)
	return nil
}

// GetPool ...
func (cs *ChainService) GetPool(args *Args, result *string) error {
	console.Dev("ChainService.GetPool()")
	B, err := json.MarshalIndent(cs.b, "", "\t")

	*result = string(B)
	return err
}

// NewChainService ...
func NewChainService() *ChainService {
	console.Dev("NewChainService()")
	CS := new(ChainService)

	path := "./chain.config"

	CS.c = new(chain.Chain)
	CS.c.Genesis(path)

	CS.b = chain.NewBlock(CS.c.BN())

	return CS
}
