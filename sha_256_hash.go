package main

import (
	"crypto/sha256"
	"encoding/hex"
)

func calculateSHA256(message string) string {

	hash := sha256.New()
	hash.Write([]byte(message))

	hashValue := hash.Sum(nil)
	hashString := hex.EncodeToString(hashValue)

	return hashString
}
