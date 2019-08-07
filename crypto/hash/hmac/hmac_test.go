package hmac

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSHA1(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")
	expectedMACBytes, _ := hex.DecodeString("de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9")
	messageMACBytes := SHA1(key, message)

	e := fmt.Sprintf("%x", expectedMACBytes)
	m := fmt.Sprintf("%x", messageMACBytes)

	if e != m {
		t.Fail()
	}
}

func TestMD5(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")
	expectedMACBytes, _ := hex.DecodeString("80070713463e7749b90c2dc24911e275")
	messageMACBytes := MD5(key, message)

	e := fmt.Sprintf("%x", expectedMACBytes)
	m := fmt.Sprintf("%x", messageMACBytes)

	if e != m {
		t.Fail()
	}
}

func TestSHA256(t *testing.T) {
	key := []byte("key")
	message := []byte("The quick brown fox jumps over the lazy dog")
	expectedMACBytes, _ := hex.DecodeString("f7bc83f430538424b13298e6aa6fb143ef4d59a14946175997479dbc2d1a3cd8")
	messageMACBytes := SHA256(key, message)

	e := fmt.Sprintf("%x", expectedMACBytes)
	m := fmt.Sprintf("%x", messageMACBytes)

	if e != m {
		t.Fail()
	}
}
