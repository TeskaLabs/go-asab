package asab

import (
	"encoding/base64"
	"math/rand"
	"os"
)

// Generate random token using Base64 characters
func GenerateBase64Token(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// os.GetEnv with default value
func GetEnv(name string, default_value string) string {
	value, found := os.LookupEnv(name)
	if !found {
		return default_value
	}
	return value
}
