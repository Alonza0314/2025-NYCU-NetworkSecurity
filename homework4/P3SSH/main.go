package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/GehirnInc/crypt/sha512_crypt"
)

const (
	ROCKYOU_PATH = "../rockyou.txt"
	OUTPUT_PATH  = "cracked.txt"
)

// 60w

var shadowEntries = []struct {
	username string
	hash     string
}{
	// {"postfix", "$6$.K3zKjOKeKDHDMi0$RUQ8PKPG082yyanbAEydSGRjaxvf6/lGyz8l3JmTAjKqrLz.2ja4/nIKpeSeFZxTEWUOW/dNLzy4DIUlq3E7P0"},
	// {"oracle", "$6$q42x7q2ut.vgn/lS$w6lmQ/EoNTTZ8tZOlrB5a4PiYnRSTrhD3hImBmwLqiezoJtF8uyWElRtny2HcX7XzoLDNm2DWQsdaBoGeARXs1"},
	{"webmaster", "$6$7ZWjjxoq5DcVJ7lZ$FbbFgsClTHidH1qjcTNNduWhemcA7imAXqyZ/5cNZfXf5alZF9JvHcOAhqFKI2HrO5nvnqKHVgd4osbfAtJ2l/"},
	{"tcpdump", "$6$Pj.F5GMgTpQqad/n$YiSEtkrhaWoetF4OPjfXbsketKZqU.6MPxrKVWJZyArW07IdcsQEE8B1c2MWK4tP7T8CpE2.AQgZA.8O8ep8Z/"},
	{"linaro", "$6$GxBkgatBm4Xs6Z1Q$StCh5Xn.jW92U5vjUgmKuAq4YcjFaeYZngCv8aRwGv0iHgeWhHm5pJ.mCK3N4y0aujJa5kbs8CUC4PJR2FMsu0"},
	// {"hplip", "$6$NmHDZ4PmEewewHmr$TREn1vMnWiKEPZGzH8ML6lAY1bcnsI/Ag7dHSzE2h/359ieFJUI98Ra40lxAFdlmGXtZMpjbiJ7ndtL8w.0p60"},
	// {"unscd", "$6$2omk7mxfxdFMNnW/$TopWFi2rkUJke9sF4u3wvdV8/guXlIJASMTw59u8m/Q65FRH1HGJVpdSQ4UEE2aCos./D/fdp/jWrfapTQJz31"},
	// {"zabbix", "$6$r3alCLBw9XsqU0fs$6mZc23axmIeB3qHKw3kiEsmttuHUPD2y6oX21qtRYpB8Y0gInI63Aja2dFP1oadqfIn9d6wfieSbk5X4MxAKz1"},
	{"omsagent", "$6$04t9Rf/AwSdW1e.t$tgt37OXgmnmIu/1Tu0sqCP0vjCPPx5iAQmNeQLtgjU.7U2fnuya9Kl9u1nE06oWdWf0X76sDeVpMETBO.GCpb0"},
	{"xpdb", "$6$kGK6d2Q8UxFXG1wD$U/TknX/DZPtbHNuosDsePp.pgU9pxGGhazCJDP7ulrDVcBy5GiVdkSjXkPh/1/GXivnVJw.qBPoI.chELFbM4/"},
}

type CrackResult struct {
	Username string
	Password string
	Hash     string
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

	fmt.Println("=== Start cracking ===")
	wg := sync.WaitGroup{}

	for userIdx, entry := range shadowEntries {
		wg.Add(1)
		go func(user string, hash string, idx int) {
			defer wg.Done()
			crypter := sha512_crypt.New()

			for j := 0; j < len(wordlist); j++ {
				if j%1e5 == 0 {
					fmt.Printf("[%s] Tried %d passwords...\n", user, j)
				}

				if err := crypter.Verify(hash, []byte(wordlist[j])); err == nil {
					ch <- fmt.Sprintf("%s:%s", user, wordlist[j])
					fmt.Printf("CRACKED! %s:%s\n", user, wordlist[j])
					return
				}
			}
			ch <- fmt.Sprintf("Failed to crack %s\n", user)
		}(entry.username, entry.hash, userIdx)
		// for i := 0; i < 5; i++ {
		// 	wg.Add(1)
		// 	go func() {
		// 		defer wg.Done()
		// 		crypter := sha512_crypt.New()
		// 		for j := i * 3e6; j < (i+1)*3e6 && j < len(wordlist); j++ {
		// 			if j%1e5 == 0 {
		// 				fmt.Printf("[%s] Tried %d passwords...\n", entry.username, j)
		// 			}
		// 			if err := crypter.Verify(entry.hash, []byte(wordlist[j])); err == nil {
		// 				ch <- fmt.Sprintf("%s:%s", entry.username, wordlist[j])
		// 				fmt.Printf("CRACKED! %s:%s\n", entry.username, wordlist[j])
		// 				return
		// 			}
		// 		}
		// 		fmt.Printf("Failed to crack %s at index %d\n", entry.username, i)
		// 	}()
		// }
		// wg.Wait()
	}
	wg.Wait()
	close(ch)
	fmt.Println("=== Cracking finished ===")
}
