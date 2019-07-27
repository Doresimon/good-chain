package chain

import (
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/Doresimon/good-chain/console"
	"github.com/Doresimon/good-chain/store"
)

const (
	BLOCKSTATE_BABY     = 0x00
	BLOCKSTATE_HEALTH   = 0x01
	BLOCKSTATE_DANGER   = 0x02
	BLOCKSTATE_UNHEALTH = 0x03
	BLOCKSTATE_DEAD     = 0x04
)

type Chain struct {
	I          int
	config     *Config
	uid        *big.Int
	blockNum   uint64
	blockState byte
	store      *store.Operator
	tiktok     *time.Ticker
}

func (this *Chain) Genesis(path string) {
	config := new(Config)
	config.read(path)
	// config.readDefault()

	this.config = config

	this.blockNum = 0
	this.blockState = BLOCKSTATE_BABY

	this.store = new(store.Operator)
	this.store.Path = "./LEVEL/" + this.config.Name
	this.store.Open()
}

func (this *Chain) WriteBlock(B *Block) {
	console.Bingo("WriteBlock " + strconv.FormatUint(this.blockNum, 16) + ". TX = " + strconv.FormatUint(uint64(len(B.Logs)), 16))

	B.BN = this.blockNum
	data, _ := json.Marshal(B)
	this.store.Write([]byte(strconv.FormatUint(this.blockNum, 16)), data)
	this.blockNum++
	B.Clear()
}

func (this *Chain) ReadBlock(BN uint64) *Block {
	B := new(Block)
	data := this.store.Read([]byte(strconv.FormatUint(BN, 16)))
	json.Unmarshal(data, B)
	return B
}

func (this *Chain) BN() uint64 {
	return this.blockNum
}

func (this *Chain) RunTicker(B *Block) {
	this.tiktok = time.NewTicker(time.Second * 10)
	go func() {
		for _ = range this.tiktok.C {
			console.Info("ticked at " + time.Now().UTC().Format(time.RFC3339))
			this.WriteBlock(B)
		}
	}()
}
func (this *Chain) StopTicker() {
	this.tiktok.Stop()
}
