package common

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

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

	tiktokOver chan int
}

// NewChainService ...
func NewChainService() *ChainService {
	console.Dev("NewChainService()")
	CS := new(ChainService)
	return CS
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

// NewLog ...
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

// CleanPool ...
func (cs *ChainService) CleanPool() {
	console.Dev("ChainService.CleanPool()")
	cs.B.Clear()
	cs.B.Number++
}

// GetBlock ...
func (cs *ChainService) GetBlock(arg uint64, result *string) error {
	console.Dev("ChainService.GetBlock()")

	B := cs.C.ReadBlock(arg)

	jsonB, err := json.MarshalIndent(B, "", "\t")

	*result = string(jsonB)
	return err
}

// RunTicker TODO
func (cs *ChainService) RunTicker() {
	tiktok := time.NewTicker(time.Millisecond * 1000 * 10)
	go func() {
		for {
			select {
			case <-tiktok.C:
				cs.C.WriteBlock(cs.B)
			case <-cs.tiktokOver:
				return
			}
		}
	}()
}

// StopTicker TODO
// func (cs *ChainService) StopTicker() {
// 	c.tiktokOver <- true
// 	c.tiktok.Stop()
// }

// HexMessage ...
type HexMessage struct {
	Pk                 string
	SupposeBlockNumber string
	Message            string
	Sig                StringSig
}

// ToLog ...
// decode hex message to specific structure
func (hm *HexMessage) ToLog() *chain.Log {
	pk, _ := hex.DecodeString(hm.Pk)
	SupposeBlockNumber, _ := hex.DecodeString(hm.SupposeBlockNumber)
	Message, _ := hex.DecodeString(hm.Message)
	// R,_ := hex.DecodeString(hm.Sig.R)
	// S,_ := hex.DecodeString(hm.Sig.S)
	// H,_ := hex.DecodeString(hm.Sig.H)
	R := hm.Sig.R
	// S := hm.Sig.S
	// H := hm.Sig.H

	// L := chain.NewLog(pk, SupposeBlockNumber, Message, R, S, H)
	L := chain.NewLog(pk, SupposeBlockNumber, Message, R)

	return L
}

// GetChain ...
// func (cs *ChainService) GetChain() *chain.Chain {
// 	return cs.C
// }

// StringSig ...
type StringSig struct {
	R []byte
	S []byte
	H []byte
}
