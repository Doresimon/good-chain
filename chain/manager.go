package chain

import (
	"encoding/json"
	"fmt"
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

// NewChain create a new chain instacne with config file path
func NewChain(path string) *Chain {
	c := new(Chain)
	c.Genesis(path)
	return c
}

// Chain TODO
type Chain struct {
	I            int
	config       *Config
	uid          *big.Int
	blockNumber  uint64
	blockState   byte
	store        *store.Operator
	tiktok       *time.Ticker
	tiktokOver   chan bool
	pendingBlock *Block
}

// Genesis TODO
func (c *Chain) Genesis(path string) {
	config := new(Config)
	config.read(path)
	// config.readDefault()

	c.config = config

	c.blockNumber = 0
	c.blockState = BLOCKSTATE_BABY

	c.store = new(store.Operator)
	c.store.Path = "./LEVEL/" + c.config.Name
	c.store.Open()
}

// WriteBlock TODO
func (c *Chain) WriteBlock(block *Block) {
	console.Bingo(fmt.Sprintf("WriteBlock block.num=0x%x tx.num=%d", c.blockNumber, len(block.Logs)))

	block.Number = c.blockNumber
	data, _ := json.Marshal(block)
	c.store.Write([]byte(strconv.FormatUint(c.blockNumber, 16)), data)
	c.blockNumber++
	block.Clear()
}

// ReadBlock TODO
func (c *Chain) ReadBlock(blockNumber uint64) *Block {
	block := new(Block)
	data := c.store.Read([]byte(strconv.FormatUint(blockNumber, 16)))
	json.Unmarshal(data, block)
	return block
}

// BN TODO
func (c *Chain) BN() uint64 {
	return c.blockNumber
}

// RunTicker TODO
func (c *Chain) RunTicker(B *Block) {
	c.tiktok = time.NewTicker(time.Millisecond * 1000 * 10)
	go func() {
		for {
			select {
			case <-c.tiktok.C:
				c.WriteBlock(B)
			case <-c.tiktokOver:
				return
			}
		}
	}()
}

// StopTicker TODO
func (c *Chain) StopTicker() {
	c.tiktokOver <- true
	c.tiktok.Stop()
}
