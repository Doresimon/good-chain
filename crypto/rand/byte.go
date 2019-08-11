package rand

import (
	"crypto/rand"
)

// Bytes ...
func Bytes(length int) ([]byte, error) {
	ret := make([]byte, length)
	_, err := rand.Read(ret)
	return ret, err
}
