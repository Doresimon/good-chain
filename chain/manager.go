package chain

import "math/big"

const (
	BLOCKSTATE_BABY     = 0x00
	BLOCKSTATE_HEALTH   = 0x01
	BLOCKSTATE_DANGER   = 0x02
	BLOCKSTATE_UNHEALTH = 0x03
	BLOCKSTATE_DEAD     = 0x04
)

type Chain struct {
	config     *ChainConfig
	uid        *big.Int
	blockNum   *big.Int
	blockState byte
}

func (this *Chain) Genesis(path string) {
	config := new(ChainConfig)
	config.read(path)
	// config.readDefault()

	this.config = config

	this.blockNum = big.NewInt(0)
	this.blockState = BLOCKSTATE_BABY
}
