package application

// Response ...
type Response struct {
	Signer   string `json:"signer"`
	Request  string `json:"request"` // the request's hash
	Data     string `json:"data"`    // main infomation
	Commit   string `json:"commit"`  // sha256(data)
	ECIESKey string `json:"ecies-key"`
}
