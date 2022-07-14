package encryption

import "testing"

func TestKeys(t *testing.T) {
	private, err := NewKeys()
	if err != nil || private == nil {
		t.Fatalf("NewKeys() failed: %v", err)
	}
}

func TestEncrypt(t *testing.T) {
	keys, _ := NewKeys()
	plaintext := []byte("Hello, world!")
	encryptedText, err := Encrypt(plaintext, &keys.PublicKey)
	if err != nil || encryptedText == nil {
		t.Fatalf("Encrypt() failed: %v", err)
	}
}

func TestDecrypt(t *testing.T) {
	keys, _ := NewKeys()
	plaintext := []byte("Hello, world!")
	encryptedText, _ := Encrypt(plaintext, &keys.PublicKey)
	decryptedText, err := Decrypt(encryptedText, keys)
	if err != nil || decryptedText == nil {
		t.Fatalf("Decrypt() failed: %v", err)
	}
}

func TestToJson(t *testing.T) {
	keys, _ := NewKeys()
	json, err := PublicKeyToJson(&keys.PublicKey)
	if err != nil || json == nil {
		t.Fatalf("PublicKeyToJson() failed: %v", err)
	}
}

func TestFromJson(t *testing.T) {
	keys, _ := NewKeys()
	json, _ := PublicKeyToJson(&keys.PublicKey)
	key, err := JsonToPublicKey(json)
	if err != nil || key == nil {
		t.Fatalf("JsonToPublicKey() failed: %v", err)
	}
}
