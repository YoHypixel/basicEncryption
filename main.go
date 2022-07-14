package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
)

func Encrypt(b []byte, key *rsa.PublicKey) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, key, b, nil)
}

func Decrypt(b []byte, key *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, key, b, nil)
}

func NewKeys() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 1024)
}

func PublicKeyToJson(key *rsa.PublicKey) ([]byte, error) {
	return json.Marshal(key)
}

func JsonToPublicKey(key []byte) (*rsa.PublicKey, error) {
	var pub *rsa.PublicKey
	err := json.Unmarshal(key, &pub)
	if err != nil {
		return nil, err
	}
	return pub, nil
}
