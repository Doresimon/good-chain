package application

import "encoding/json"

// Response ...
type Response struct {
	Responser      *Account
	ResponserPath  string
	RequestHashHex string
	Data           string
	ECIESKey       string
	TXHash         string
}

// ResponseCreation ...
type ResponseCreation struct {
	Signer   string `json:"signer"`
	Request  string `json:"request"` // the request's hash
	Data     string `json:"data"`    // main infomation
	Commit   string `json:"commit"`  // sha256(data)
	ECIESKey string `json:"ecies-key"`
}

// ParseResponseCreation ...
func ParseResponseCreation(contentBytes []byte) *Response {
	rc := new(ResponseCreation)
	err := json.Unmarshal(contentBytes, rc)
	if err != nil {
		panic(err)
	}

	res := new(Response)
	res.Responser = nil
	res.ResponserPath = rc.Signer
	res.RequestHashHex = rc.Request
	res.Data = rc.Data
	res.ECIESKey = rc.ECIESKey
	return res
}
