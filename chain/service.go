package chain

import (
	"encoding/json"
	"time"

	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/crypto/bls"
	"golang.org/x/crypto/bn256"
)

// LogTransferPool is used to transfer data with p2p service
var LogTransferPool = make(chan *Log, 256)

// NewService ...
func NewService(c *Chain) *Service {
	console.Devf("chain.NewService() >> logPoolSize = %d, blockPoolSize = %d", logPoolSize, blockPoolSize)

	cs := new(Service)
	cs.LogPool = make(chan *Log, logPoolSize)
	cs.BlockPool = make(chan *Block, blockPoolSize)
	cs.C = c

	go cs.MoniterPool()
	go cs.RunTicker()

	return cs
}

// Service ...
type Service struct {
	C *Chain

	// LogPool   []*Log
	// BlockPool []*Block
	LogPool   chan *Log
	BlockPool chan *Block

	tiktokOver chan interface{}
}

// AddLog add a log to LogPool
func (cs *Service) AddLog(log *Log) {
	cs.LogPool <- log
}

// AddBlock add a log to blockPool
func (cs *Service) AddBlock(block *Block) {
	cs.BlockPool <- block
}

// ProduceBlock produce a new block with current logs in pool
func (cs *Service) ProduceBlock() *Block {
	console.Infof("chain.Service.ProduceBlock() >> log.Len = %d", len(cs.LogPool))

	logLen := len(cs.LogPool)
	b := NewBlock(cs.C.blockNumber)
	for i := 0; i < logLen; i++ {
		b.AddLog(<-cs.LogPool)
	}
	return b
}

// RunTicker TODO
func (cs *Service) RunTicker() {
	tiktok := time.NewTicker(blockTime)
	for {
		select {
		case <-tiktok.C:
			b := cs.ProduceBlock()
			cs.C.WriteBlock(b)
		case <-cs.tiktokOver:
			tiktok.Stop()
			return
		}
	}
}

// StopTicker TODO
func (cs *Service) StopTicker() {
	cs.tiktokOver <- true
}

// MoniterPool ...
func (cs *Service) MoniterPool() {
	for {
		var l = new(Log)
		l = <-LogTransferPool

		Sender := l.Sender
		SupposeBN := l.SupposeBN
		TX := l.TX
		Sig := l.Sig

		console.Warnf("Sender: %s", Sender)
		console.Warnf("SupposeBN: %s", SupposeBN)
		console.Warnf("TX: %s", TX)
		console.Warnf("Sig: %s", Sig)

		pk := new(bn256.G2)
		pk.Unmarshal(Sender)
		sig := new(bn256.G1)
		sig.Unmarshal(Sig)

		ok := bls.Verify(pk, TX, sig)
		if ok {
			console.Bingo("verufy sig success")
		} else {
			console.Error("verufy sig fail")
		}

		cs.AddLog(l)
	}
}

// CreateBlock ...
// func (cs *Service) CreateBlock(args *Args, result *string) error {
// 	console.Dev("Service.CreateBlock()")

// 	for s := range args.Data {
// 		fmt.Println(s)
// 	}

// 	*result = "hiahiahia"
// 	return nil
// }

// // NewLog ...
// func (cs *Service) NewLog(args *HexMessage, result *string) error {
// 	cs.I++
// 	console.Dev("Service.NewLog()" + strconv.Itoa(cs.I))

// 	L := args.ToLog()

// 	// ok := L.VerifySig()
// 	ok := true

// 	if !ok {
// 		console.Error("Signature verify failed")
// 		return errors.New("Signature verify failed")
// 	} else {
// 		console.Bingo("Signature verify success")
// 	}

// 	cs.B.AddLog(L)

// 	Ls, _ := json.Marshal(L)

// 	*result = string(Ls)
// 	return nil
// }

// // GetPool ...
// func (cs *Service) GetPool(args *Args, result *string) error {
// 	console.Dev("Service.GetPool()")
// 	B, err := json.MarshalIndent(cs.B, "", "\t")

// 	*result = string(B)
// 	return err
// }

// // CleanPool ...
// func (cs *Service) CleanPool() {
// 	console.Dev("Service.CleanPool()")
// 	cs.B.Clear()
// 	cs.B.Number++
// }

// GetBlock ...
func (cs *Service) GetBlock(arg uint64, result *string) error {
	console.Dev("Service.GetBlock()")

	B := cs.C.ReadBlock(arg)

	jsonB, err := json.MarshalIndent(B, "", "\t")

	*result = string(jsonB)
	return err
}
