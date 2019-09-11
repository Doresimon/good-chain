package chain

// NewBlock create a new block instance
func NewBlock(BN uint64) *Block {
	b := new(Block)
	b.Logs = make([]Log, 0, 0)
	return b
}

// Block ...
type Block struct {
	Number uint64 `json:"block number"`
	Logs   []Log  `json:"logs"`
}

// AddLog adds a new log to a block
func (block *Block) AddLog(L *Log) error {
	block.Logs = append(block.Logs, *L)
	return nil
}

// Clear clears a block's logs
func (block *Block) Clear() {
	block.Logs = make([]Log, 0, 0)
}
