package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
)

// MD5 hmac-md5
func MD5(key []byte, message []byte) []byte {
	mac := hmac.New(md5.New, key)
	mac.Write(message)
	messageMAC := mac.Sum(nil)
	return messageMAC
}

// SHA1 hmac-sha1
func SHA1(key []byte, message []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	messageMAC := mac.Sum(nil)
	return messageMAC
}

// SHA256 hmac-sha256
func SHA256(key []byte, message []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	messageMAC := mac.Sum(nil)
	return messageMAC
}

// SHA512 hmac-sha512
func SHA512(key []byte, message []byte) []byte {
	mac := hmac.New(sha512.New, key)
	mac.Write(message)
	messageMAC := mac.Sum(nil)
	return messageMAC
}
