package types

const (
	RootAccountCreation = iota // 创建组织账号, 主要为添加信息到账号中
)

// Transaction .
type Transaction struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
	KeyPath   string `json:"key-path"`  // pk/0/2/9
	TimeStamp string `json:"timestamp"` // milliseconds
	Nonce     string `json:"nonce"`     // an unique number of an account
	// PublicKey          string `json:"public-key"`
	// SignatureChain     string `json:"signature-chain"`
	// AggregateSignature string `json:"aggregate-signature"`
}

// Sign .
func (tx *Transaction) Sign() {
	return
}
