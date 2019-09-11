package application

// OrgCreation is used to parse a message
type OrgCreation struct {
	Name      string `json:"name"`
	PublicKey string `json:"pubkey"`
	ChainCode string `json:"chaincode"`
	Extra     string `json:"extra"`
}
