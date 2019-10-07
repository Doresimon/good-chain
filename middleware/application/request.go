package application

import "encoding/json"

// Request ...
type Request struct {
	Requester     *Account
	RequesterPath string
	Data          string
	ECIESKey      string
	TXHash        string
}

// RequestCreation ...
type RequestCreation struct {
	Signer   string `json:"signer"`
	Data     string `json:"data"`
	Commit   string `json:"commit"`
	ECIESKey string `json:"ecies-key"`
}

// ParseRequestCreation ...
func ParseRequestCreation(contentBytes []byte) *Request {
	rc := new(RequestCreation)
	err := json.Unmarshal(contentBytes, rc)
	if err != nil {
		panic(err)
	}

	req := new(Request)
	req.RequesterPath = rc.Signer
	req.Data = rc.Data
	req.ECIESKey = rc.ECIESKey
	req.Requester = nil
	return req
}
