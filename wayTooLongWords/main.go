package main

import (
	"fmt"
	"strconv"
)

func main() {
	words := readInput()
	abbreviate(words)
}

func abbreviate(words []string) {
	for i := 0; i < len(words); i++ {
		if len(words[i]) > 10 {
			lenWord := len(words[i]) - 2
			newWord := string(words[i][0])
			newWord += strconv.Itoa(lenWord)
			newWord += string(words[i][len(words[i])-1])
			fmt.Println(newWord)
		} else {
			fmt.Println(words[i])
		}
	}
}

func readInput() []string {
	var nWords int
	var words []string
	var word string

	fmt.Scanln(&nWords)

	for i := 0; i < nWords; i++ {
		fmt.Scanln(&word)
		words = append(words, word)
	}

	return words
}
