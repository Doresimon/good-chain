package state

import (
	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/middleware/application"
)

// Pool ...
type Pool struct {
	accounts map[string]*application.Account
	txs      map[string]chain.Body
}
