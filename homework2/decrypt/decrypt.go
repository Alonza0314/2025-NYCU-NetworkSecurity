package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"hw2/util"
)

var (
	aesKey []byte
)

func Decrypt(ciphertextFile string, plaintextFile string, key string) error {
	if !util.IsFileExists(ciphertextFile) {
		return fmt.Errorf("ciphertext file does not exist")
	}

	ciphertext, err := util.ReadFile(ciphertextFile)
	if err != nil {
		return fmt.Errorf("failed to read ciphertext file: %v", err)
	}
	fmt.Printf("Read ciphertext file, content length: %d bytes\n", len(ciphertext))

	if len(key) == 0 {
		return fmt.Errorf("decryption key must be provided")
	}

	aesKey = make([]byte, 32)
	n, err := fmt.Sscanf(key, "%x", &aesKey)
	if err != nil || n != 1 {
		return fmt.Errorf("invalid key format. Must be 64 hex characters (32 bytes)")
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %v", err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		return fmt.Errorf("ciphertext too short")
	}
	iv, ciphertext := ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("failed to decrypt: %v", err)
	}
	fmt.Printf("Decrypted content length: %d bytes\n", len(plaintext))

	if err := util.WriteFile(plaintextFile, plaintext); err != nil {
		return fmt.Errorf("failed to write plaintext file: %v", err)
	}

	fmt.Printf("File decrypted successfully: %s\n", plaintextFile)
	return nil
}
