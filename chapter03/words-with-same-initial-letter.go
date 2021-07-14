package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	occurrences := getOccurrences(words)

	print(occurrences)
}

func getOccurrences(words []string) map[string]int {
	occurrences := make(map[string]int)

	for _, word := range words {
		initial := strings.ToUpper(string(word[0]))
		counter, found := occurrences[initial]

		if found {
			occurrences[initial] = counter + 1
		} else {
			occurrences[initial] = 1
		}
	}

	return occurrences
}

func print(occurrences map[string]int) {
	fmt.Println("Word count starting with each letter")

	for initial, occurrence := range occurrences {
		fmt.Printf("%s = %d \n", initial, occurrence)
	}
}
