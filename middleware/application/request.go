package application

// Request ...
type Request struct {
	Signer   string `json:"signer"`
	Data     string `json:"data"`
	Commit   string `json:"commit"`
	ECIESKey string `json:"ecies-key"`
}
