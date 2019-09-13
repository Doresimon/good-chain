package chain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Body ...
type Body struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	// Content      *Content `json:"content"`
	ContentBytes []byte `json:"content-bytes"`
	Timestamp    uint32 `json:"timestamp"`
}

// Hash ...
func (b *Body) Hash() []byte {
	bodyBytes, _ := json.Marshal(b)
	hash := sha256.Sum256(bodyBytes)
	return hash[:]
}

// HashHexString ...
func (b *Body) HashHexString() string {
	return fmt.Sprintf("%x", b.Hash())
}

// Content ...
// type Content struct {
// 	Name  string            `json:"name"`
// 	Path  string            `json:"path"`
// 	Index string            `json:"index"`
// 	Pk    string            `json:"pk"`
// 	Extra map[string]string `json:"extra"`
// }
