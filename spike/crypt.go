package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

// VerifyMessage verifies if a message was encrypted with the given key
// and produced the actual digest
func VerifyMessage(message, key string, actualDigest []byte) bool {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(message))
	expectedMac := mac.Sum(nil)
	fmt.Println(expectedMac)
	fmt.Println(actualDigest)
	return hmac.Equal(expectedMac, actualDigest)
}
