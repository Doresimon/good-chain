package types

// Transaction .
type Transaction struct {
	Type               string `json:"type"`
	Content            string `json:"content"`
	PublicKey          string `json:"public-key"`
	SignatureChain     string `json:"signature-chain"`
	AggregateSignature string `json:"aggregate-signature"`
	SenderHDPath       string `json:"sender-HD-path"` // HD path of sender
}

// Sign .
func (tx *Transaction) Sign() {
	return
}
