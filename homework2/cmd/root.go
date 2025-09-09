package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fileCrypt",
	Short: "A file encryption and decryption tool using AES-256-GCM",
	Long: `FileCrypt is a secure file encryption and decryption tool that uses AES-256-GCM.
	
It supports:
- AES-256 encryption with GCM mode for authenticated encryption
- Random key generation or user-provided keys
- Secure IV (nonce) generation for each encryption
- File-based input and output`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
