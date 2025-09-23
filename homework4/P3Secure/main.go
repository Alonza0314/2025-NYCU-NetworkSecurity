package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"slices"
	"strings"
	"sync"

	"golang.org/x/crypto/sha3"
)

const (
	ROCKYOU_PATH = "../rockyou.txt"
	OUTPUT_PATH  = "cracked.txt"
)

var md5Hashes = []string{
	// "e49201c3a8f548902b9ae9f16638f879",
	// "19cf9dda4107b300d3218702df95c76d",
	// "c6281df39e8ade06c6cc9e0095fd5c0f",
	// "a54034981409ed58d584dc9051853ddb",
	// "f58291f81868320f11235d9b9d416115",
	// "ce1c96461fbb2ad92fffcafafe85d0d1",
	// "c6177167ebb2c37352c3a63f6fa0c39d",
	// "5993428babd2cb253834e06de1800916",
	// "bebc51b6f0bbd5da67950200a89026f6",
	// "456c5a41af2eb09ac0ba0eb64f614887",
}

var shaHashes = []string{
	// "1074f17769cc2dfc0d65f713a7d8c4fd97fc78c69cfa13263b07b0e40b3cf83a",
	// "94f72dc2ea6bfae657b0ee3d5adb992aa669f6c4141717344e24e873dc09be04",
	// "19c743dc300d52fc93b5ee8c6d224f3beb8a05079e6439855cdae7e55bf16ef0",
	// "20e5b0556c431db9a147c3f73a0ae03d12f5ef391d277cd59ff0f2dd98198ec5",
	// "a44cf105063b06bbb160c22058e9c3137c8ef424ae72f981d73b10fdc743026f",
	// "74151544815c4a0153c2e7dfabcfd066d510d6996148d6c02f246c9c497bd15c",
	"745af7302284f80ddadf6893f64e247334aa899bfe90512a59aa41ea2863f56a",
	"9d34ebe967a790ada61cfa2b4e16671bfb18f0ff59296f24a0eec20dacc5ece3",
	// "0ecd9ac47c8e4b059c2b97db9657f80f203454ac8fcb01976e1decdb30af2510",
	// "3b2918324171f88304baee77d71cc0abd40e12f16f9a22404736000f00a7c7b6",
}

func outputCracked(ch chan string, ctx context.Context) {
	file, err := os.OpenFile(OUTPUT_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		return
	}
	defer file.Close()

	for {
		select {
		case output, ok := <-ch:
			if !ok {
				return
			}
			file.WriteString(output + "\n")
		case <-ctx.Done():
			return
		}
	}
}

func loadWordList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	return words, scanner.Err()
}

func md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func sha3Hash(text string) string {
	hash := sha3.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	fmt.Println("=== Loading wordlist ===")
	wordlist, err := loadWordList(ROCKYOU_PATH)
	if err != nil {
		fmt.Printf("Error loading wordList: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d passwords\n", len(wordlist))

	fmt.Println("=== Init output writer ===")
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go outputCracked(ch, ctx)

	fmt.Println("=== Starting hash cracking ===")
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for j := i * 4e6; j < (i+1)*4e6 && j < len(wordlist); j++ {
				if j%1e5 == 0 {
					fmt.Printf("Trying %d passwords...\n", j)
				}
				md5hash := md5Hash(wordlist[j])
				sha3hash := sha3Hash(wordlist[j])
				if slices.Contains(md5Hashes, md5hash) {
					ch <- fmt.Sprintf("MD5: %s: %s", md5hash, wordlist[j])
					fmt.Printf("MD5: %s: %s\n", md5hash, wordlist[j])
				}
				if slices.Contains(shaHashes, sha3hash) {
					ch <- fmt.Sprintf("SHA3: %s: %s", sha3hash, wordlist[j])
					fmt.Printf("SHA3: %s: %s\n", sha3hash, wordlist[j])
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	fmt.Println("=== Cracking finished ===")
}
