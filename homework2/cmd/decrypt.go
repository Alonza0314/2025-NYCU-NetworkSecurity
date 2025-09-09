package cmd

import (
	"fmt"
	"hw2/decrypt"
	"os"

	"github.com/spf13/cobra"
)

var (
	decryptInput  string
	decryptOutput string
	decryptKey    string
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt a file using AES-256",
	Run:   decryptFunc,
}

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&decryptInput, "input", "i", "", "Input file to decrypt")
	decryptCmd.Flags().StringVarP(&decryptOutput, "output", "o", "", "Output file for decrypted data")
	decryptCmd.Flags().StringVarP(&decryptKey, "key", "k", "", "Decryption key (hex format, 32 bytes / 64 hex chars)")

	if err := decryptCmd.MarkFlagRequired("input"); err != nil {
		fmt.Println("Error marking input flag as required:", err)
		os.Exit(1)
	}

	if err := decryptCmd.MarkFlagRequired("output"); err != nil {
		fmt.Println("Error marking output flag as required:", err)
		os.Exit(1)
	}
	decryptCmd.MarkFlagRequired("output")
	if err := decryptCmd.MarkFlagRequired("key"); err != nil {
		fmt.Println("Error marking key flag as required:", err)
		os.Exit(1)
	}
}

func decryptFunc(cmd *cobra.Command, args []string) {
	if err := decrypt.Decrypt(decryptInput, decryptOutput, decryptKey); err != nil {
		fmt.Println("Error decrypting file:", err)
		os.Exit(1)
	}
}
