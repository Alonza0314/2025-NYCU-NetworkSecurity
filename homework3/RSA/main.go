package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("Welcome to use NYCU RSAer.")
	fmt.Println("0. Run RSA Keypair Generator")
	fmt.Println("1. Run RSA encryptor")
	fmt.Println("2. Run RSA decryptor")
	fmt.Print("Enter your choice: ")

	choice, _ := reader.ReadString('\n')
	if len(choice) > 0 {
		switch choice[0] {
		case '0':
			keypair()
		case '1':
			encryptor()
		case '2':
			decryptor()
		default:
			fmt.Println("Unknown choice.")
			os.Exit(0)
		}
	}
}

func keypair() {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Failed to generate key pair: %v\n", err)
		return
	}

	// Convert private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	privateKeyFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Printf("Failed to create private key file: %v\n", err)
		return
	}
	pem.Encode(privateKeyFile, privateKeyPEM)
	privateKeyFile.Close()

	// Convert public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}
	publicKeyFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Printf("Failed to create public key file: %v\n", err)
		return
	}
	pem.Encode(publicKeyFile, publicKeyPEM)
	publicKeyFile.Close()

	fmt.Println("Key pair generated:")
	fmt.Println("- private.pem (Private Key)")
	fmt.Println("- public.pem (Public Key)")
}

func encryptor() {
	// Read public key file
	publicKeyData, err := os.ReadFile("public.pem")
	if err != nil {
		fmt.Printf("Failed to read public key file: %v\n", err)
		return
	}

	block, _ := pem.Decode(publicKeyData)
	if block == nil {
		fmt.Println("Failed to parse public key")
		return
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse public key: %v\n", err)
		return
	}

	fmt.Print("Enter text to encrypt: ")
	message, _ := reader.ReadString('\n')
	message = message[:len(message)-1] // Remove newline

	// Encrypt data
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		fmt.Printf("Encryption failed: %v\n", err)
		return
	}

	// Write encrypted data to file
	err = os.WriteFile("encrypted.txt", ciphertext, 0644)
	if err != nil {
		fmt.Printf("Failed to write encrypted file: %v\n", err)
		return
	}

	fmt.Println("Encryption complete, saved to encrypted.txt")
}

func decryptor() {
	// Read private key file
	privateKeyData, err := os.ReadFile("private.pem")
	if err != nil {
		fmt.Printf("Failed to read private key file: %v\n", err)
		return
	}

	block, _ := pem.Decode(privateKeyData)
	if block == nil {
		fmt.Println("Failed to parse private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse private key: %v\n", err)
		return
	}

	// Read encrypted file
	ciphertext, err := os.ReadFile("encrypted.txt")
	if err != nil {
		fmt.Printf("Failed to read encrypted file: %v\n", err)
		return
	}

	// Decrypt data
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		fmt.Printf("Decryption failed: %v\n", err)
		return
	}

	fmt.Printf("Decryption result: %s\n", string(plaintext))
}
