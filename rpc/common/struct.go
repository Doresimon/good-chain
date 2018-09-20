package common

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"good-chain/chain"
	"good-chain/console"
	"strconv"
)

// Args ...
// ajax data
type Args struct {
	Data []string
}

// ChainService ...
type ChainService struct {
	I int
	c *chain.Chain
	b *chain.Block
}

// HexMessage ...
type HexMessage struct {
	Pk                 string
	SupposeBlockNumber string
	Message            string
	Sig                StringSig
}

// HexMessage ...
type StringSig struct {
	R []byte
	S []byte
	H []byte
}

// ToLog ...
// decode hex message to specific structure
func (this *HexMessage) ToLog() *chain.Log {
	pk, _ := hex.DecodeString(this.Pk)
	SupposeBlockNumber, _ := hex.DecodeString(this.SupposeBlockNumber)
	Message, _ := hex.DecodeString(this.Message)
	// R,_ := hex.DecodeString(this.Sig.R)
	// S,_ := hex.DecodeString(this.Sig.S)
	// H,_ := hex.DecodeString(this.Sig.H)
	R := this.Sig.R
	S := this.Sig.S
	H := this.Sig.H

	L := chain.NewLog(pk, SupposeBlockNumber, Message, R, S, H)

	return L
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

func (cs *ChainService) NewLog(args *HexMessage, result *string) error {
	cs.I++
	console.Dev("ChainService.NewLog()" + strconv.Itoa(cs.I))

	L := args.ToLog()

	ok := L.VerifySig()

	if !ok {
		console.Error("Signature verify failed")
		return errors.New("Signature verify failed")
	} else {
		console.Bingo("Signature verify success")
	}

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

// GetChain ...
func (cs *ChainService) GetChain() *chain.Chain {
	return cs.c
}

// NewChainService ...
func NewChainService() *ChainService {
	console.Dev("NewChainService()")
	CS := new(ChainService)

	path := "./chain.config"

	CS.c = new(chain.Chain)
	CS.c.Genesis(path)
	CS.c.I = 0

	CS.b = chain.NewBlock(CS.c.BN())

	return CS
}
