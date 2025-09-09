package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"hw2/util"
	"io"
)

var (
	aesKey []byte
)

func Encrypt(plaintextFile string, ciphertextFile string, key string) error {
	if !util.IsFileExists(plaintextFile) {
		return fmt.Errorf("plaintext file does not exist")
	}

	plaintext, err := util.ReadFile(plaintextFile)
	if err != nil {
		return fmt.Errorf("failed to read plaintext file: %v", err)
	}
	fmt.Printf("Read plaintext file, content length: %d bytes\n", len(plaintext))

	if len(key) == 0 {
		fmt.Println("No key provided, generating a random key")
		aesKey = make([]byte, 32)
		if _, err := io.ReadFull(rand.Reader, aesKey); err != nil {
			return fmt.Errorf("failed to generate key: %v", err)
		}
		fmt.Printf("Generated key (save this for decryption): %x\n", aesKey)
		key = fmt.Sprintf("%x", aesKey)
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %v", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	if err := util.WriteFile(ciphertextFile, ciphertext); err != nil {
		return fmt.Errorf("failed to write ciphertext file: %v", err)
	}

	fmt.Printf("File encrypted successfully: %s\n", ciphertextFile)
	return nil
}
