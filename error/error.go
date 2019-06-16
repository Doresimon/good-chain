package error

import "github.com/Doresimon/good-chain/console"

// Check ...
// check if any error shows up
func Check(info string, err error) {
	if err != nil {
		console.Error(info + err.Error())
	}
}
