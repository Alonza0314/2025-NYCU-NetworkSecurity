package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/GehirnInc/crypt/apr1_crypt"
)

const (
	ROCKYOU_PATH = "../rockyou.txt"
	HTPASSWD_PATH = "htpasswd.txt"
	OUTPUT_PATH   = "cracked.txt"
)

type htpasswdEntry struct {
	username string
	hash     string
	salt     string
	expected string
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

func parseHtpasswd(filename string) ([]htpasswdEntry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []htpasswdEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		username, hash := parts[0], parts[1]

		hashParts := strings.Split(hash, "$")
		entry := htpasswdEntry{
			username: username,
			hash:     hash,
			salt:     hashParts[2],
			expected: hashParts[3],
		}
		entries = append(entries, entry)
	}

	return entries, scanner.Err()
}

func crackPassword(entry htpasswdEntry, wordlist []string, ch chan string, wg *sync.WaitGroup) {
	fmt.Printf("Start cracking password for %s\n", entry.username)
	defer wg.Done()

	crypter := apr1_crypt.New()

	for i, password := range wordlist {
		if i%1000 == 0 && i > 0 {
			fmt.Printf("[%s] Tried %d passwords...\n", entry.username, i)
		}

		// Use the library to verify the password
		err := crypter.Verify(entry.hash, []byte(password))
		if err == nil {
			ch <- fmt.Sprintf("%s:%s", entry.username, password)
			return
		}
	}

	ch <- fmt.Sprintf("[%s] Failed to crack password", entry.username)
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

func main() {
	fmt.Println("=== Loading htpasswd ===")
	entries, err := parseHtpasswd(HTPASSWD_PATH)
	if err != nil {
		fmt.Printf("Error loading htpasswd: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d entries\n", len(entries))

	fmt.Println("=== Loading wordList ===")
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

	fmt.Println("=== Start cracking ===")
	wg := sync.WaitGroup{}
	for _, entry := range entries {
		wg.Add(1)
		go crackPassword(entry, wordlist, ch, &wg)
	}
	wg.Wait()
	close(ch)
	fmt.Println("=== Cracking finished ===")
}
