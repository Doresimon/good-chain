package common

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/console"
)

// Args ...
// ajax data
type Args struct {
	Data []string
}

// ChainService ...
type ChainService struct {
	I int
	C *chain.Chain
	B *chain.Block
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
	// S := this.Sig.S
	// H := this.Sig.H

	// L := chain.NewLog(pk, SupposeBlockNumber, Message, R, S, H)
	L := chain.NewLog(pk, SupposeBlockNumber, Message, R)

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

	// ok := L.VerifySig()
	ok := true

	if !ok {
		console.Error("Signature verify failed")
		return errors.New("Signature verify failed")
	} else {
		console.Bingo("Signature verify success")
	}

	cs.B.AddLog(L)

	Ls, _ := json.Marshal(L)

	*result = string(Ls)
	return nil
}

// GetPool ...
func (cs *ChainService) GetPool(args *Args, result *string) error {
	console.Dev("ChainService.GetPool()")
	B, err := json.MarshalIndent(cs.B, "", "\t")

	*result = string(B)
	return err
}

// GetPool ...
func (cs *ChainService) GetBlock(arg uint64, result *string) error {
	console.Dev("ChainService.GetBlock()")

	B := cs.C.ReadBlock(arg)

	jsonB, err := json.MarshalIndent(B, "", "\t")

	*result = string(jsonB)
	return err
}

// GetChain ...
// func (cs *ChainService) GetChain() *chain.Chain {
// 	return cs.C
// }

// NewChainService ...
func NewChainService() *ChainService {
	console.Dev("NewChainService()")
	CS := new(ChainService)
	return CS
}
