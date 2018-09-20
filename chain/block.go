package chain

// Block ...
type Block struct {
	BN   uint64 `json:"block number"`
	Logs []Log  `json:"logs"`
}

func (this *Block) AddLog(L *Log) error {
	this.Logs = append(this.Logs, *L)
	return nil
}

func (this *Block) Clear() {
	this.Logs = make([]Log, 0, 0)
}

func NewBlock(BN uint64) *Block {
	B := new(Block)

	B.Logs = make([]Log, 0, 0)

	return B
}
