package message

import (
	"encoding/json"
	"fmt"

	"github.com/Doresimon/good-chain/chain"
	"github.com/Doresimon/good-chain/crypto/hdk"
)

// Parse ...
func Parse(msg []byte) {
	log, err := chain.UnmarshalLog(msg)
	if err != nil {
		panic(err)
	}
	msgBytes, err := json.Marshal(log.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("log.Sender: %x\n", log.Sender)
	fmt.Printf("msgBytes: %x\n", msgBytes)
	fmt.Printf("log.Sig: %x\n", log.Sig)

	// verify signature
	ok := hdk.Verify(log.Sender, msgBytes, log.Sig)

	if !ok {
		panic("sig verify failed")
	}

	chain.LogTransferPool <- log // chainService will monitor this channel
}
