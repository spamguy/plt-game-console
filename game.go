package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomCapitalLetters(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := make([]byte, n)
	for i := 0; i < n; i++ {
		letters[i] = byte(rand.Intn(26) + 'A') // 'A' = 65
	}
	return string(letters)
}

func main() {
	fmt.Println(randomCapitalLetters(3))
}

