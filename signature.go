package topazsdk

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func HashSum(i []byte) []byte {
	hash := sha256.New()
	hash.Write(i)
	return hash.Sum(nil)
}

func Sign(msg string, privateKey rsa.PrivateKey) string {
	msgSum := HashSum([]byte(msg))
	signature, err := rsa.SignPSS(rand.Reader, &privateKey, crypto.SHA256, msgSum, nil)
	if err != nil {
		return ""
	} else {
		return base64.StdEncoding.EncodeToString(signature)
	}
}
