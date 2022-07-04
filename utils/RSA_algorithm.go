package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

var secretKey = "m=p4tR}~yng[7+$"
var bitSize = 1024
var PrivateKey, GenerateKeyError = rsa.GenerateKey(rand.Reader, bitSize)

func CheckError(err error) error {
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func RSAOAEPEncrypt(secretMessage string, key rsa.PublicKey) string {
	label := []byte(secretKey)
	rng := rand.Reader
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(secretMessage), label)
	err = CheckError(err)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func RSAOAEPDecrypt(cipherText string, privateKey rsa.PrivateKey) string {
	ct, _ := base64.StdEncoding.DecodeString(cipherText)
	label := []byte(secretKey)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, &privateKey, ct, label)
	err = CheckError(err)
	if err != nil {
		return ""
	}
	return string(plaintext)
}
