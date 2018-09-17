package common

import (
	"fmt"
)

// Args ...
// ajax data
type Args struct {
	Data []string
}

// ChainService ...
type ChainService byte

// CreateBlock ...
func (cs *ChainService) CreateBlock(args *Args, result *string) error {
	fmt.Println("[CALL] ChainService.CreateBlock()")

	for s := range args.Data {
		fmt.Println(s)
	}

	*result = "hiahiahia"
	return nil
}
