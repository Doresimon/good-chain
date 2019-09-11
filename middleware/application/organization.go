package application

import "encoding/json"

type OrgCreation struct {
	Name      string `json:"name"`
	PublicKey string `json:"pubkey"`
	ChainCode string `json:"chaincode"`
	Extra     string `json:"extra"`
}

func HandleOrg(contentBytes []byte) {
	oc := new(OrgCreation)
	err := json.Unmarshal(contentBytes, oc)
	if err != nil {
		panic(err)
	}

}
