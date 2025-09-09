package cmd

import (
	"fmt"
	"hw2/encrypt"
	"os"

	"github.com/spf13/cobra"
)

var (
	encryptInput  string
	encryptOutput string
	keyString     string
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt a file using AES-256",
	Run:   encryptFunc,
}

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().StringVarP(&encryptInput, "input", "i", "", "Input file to encrypt")
	encryptCmd.Flags().StringVarP(&encryptOutput, "output", "o", "", "Output file for encrypted data")
	encryptCmd.Flags().StringVarP(&keyString, "key", "k", "", "Encryption key (hex format, 32 bytes / 64 hex chars). If not provided, a random key will be generated")

	if err := encryptCmd.MarkFlagRequired("input"); err != nil {
		fmt.Println("Error marking input flag as required:", err)
		os.Exit(1)
	}

	if err := encryptCmd.MarkFlagRequired("output"); err != nil {
		fmt.Println("Error marking output flag as required:", err)
		os.Exit(1)
	}
}

func encryptFunc(cmd *cobra.Command, args []string) {
	if err := encrypt.Encrypt(encryptInput, encryptOutput, keyString); err != nil {
		fmt.Println("Error encrypting file:", err)
		os.Exit(1)
	}
}
