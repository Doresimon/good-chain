package chain

import (
	"encoding/json"
	"good-chain/db"
	"math/big"
)

const (
	BLOCKSTATE_BABY     = 0x00
	BLOCKSTATE_HEALTH   = 0x01
	BLOCKSTATE_DANGER   = 0x02
	BLOCKSTATE_UNHEALTH = 0x03
	BLOCKSTATE_DEAD     = 0x04
)

type Chain struct {
	config     *Config
	uid        *big.Int
	blockNum   uint64
	blockState byte
	db         *db.Operator
}

func (this *Chain) Genesis(path string) {
	config := new(Config)
	config.read(path)
	// config.readDefault()

	this.config = config

	this.blockNum = 0
	this.blockState = BLOCKSTATE_BABY

	this.db = new(db.Operator)
	this.db.Path = "./LEVEL/" + this.config.Name
	this.db.Open()
}

func (this *Chain) WriteBlock(B *Block) {
	B.BN = this.blockNum
	data, _ := json.Marshal(B)
	this.db.Write([]byte(string(this.blockNum)), data)
}

func (this *Chain) ReadBlock(BN uint64) *Block {
	B := new(Block)
	data := this.db.Read([]byte(string(BN)))
	json.Unmarshal(data, B)
	return B
}

func (this *Chain) BN() uint64 {
	return this.blockNum
}
