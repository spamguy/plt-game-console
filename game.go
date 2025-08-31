package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"slices"
	"time"
)

func randomCapitalLetters(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := make([]byte, n)
	for i := 0; i < n; i++ {
		letters[i] = byte(r.Intn(26) + 'A') // 'A' = 65
	}
	return string(letters)
}

func main() {
	file, err := os.Open("./english-words/words_alpha.txt")

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close() // Ensure the file is closed
	scanner := bufio.NewScanner(file)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	var wordInput string
	var startTime time.Time
	var words []string
	var betterWords []string

	clue := randomCapitalLetters(3)
	searchExpression := fmt.Sprintf("(?i)%c[a-z]*%c[a-z]*%c", clue[0], clue[1], clue[2])
	isInputMatch := false
	isRealWord := false

	for !isRealWord {
		for !isInputMatch {
			fmt.Println(clue)
			fmt.Scanln(&wordInput)

			if isInputMatch, _ = regexp.MatchString(searchExpression, wordInput); !isInputMatch {
				fmt.Println(clue, "isn't present in", wordInput)
			}
		}

		startTime = time.Now()

		for scanner.Scan() {
			w := scanner.Text()
			words = append(words, w)

			if isMatch, _ := regexp.MatchString(searchExpression, w); isMatch && len(w) <= len(wordInput) {
				betterWords = append(betterWords, w)
			}
		}

		if isRealWord = slices.Contains(betterWords, wordInput); !isRealWord {
			fmt.Println(wordInput, "isn't a real word; try again :(")
			isInputMatch = false
		}
	}

	fmt.Println(len(words), "words read in", time.Since(startTime))

	fmt.Println("Equal or better words:", betterWords)
}
