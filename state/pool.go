package state

import (
	"github.com/Doresimon/good-chain/chain"
)

// Pool ...
type Pool struct {
	accounts map[string]*Account
	txs      map[string]chain.Body
}
