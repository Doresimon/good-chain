package chain

// Log ...
type Log struct {
	Sender    string `json:"sender"`
	SupposeBN string `json:"suppose Block number"`
	Message   string `json:"message"`
}

func NewLog(s string, sBN string, m string) *Log {
	L := new(Log)
	L.Sender = s
	L.SupposeBN = sBN
	L.Message = m

	return L
}
