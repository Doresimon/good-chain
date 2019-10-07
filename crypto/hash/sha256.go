package hash

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 ...
func SHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// SHA256Hex ...
func SHA256Hex(data []byte) string {
	return fmt.Sprintf("%x", SHA256(data))
}
