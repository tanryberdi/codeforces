package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	st, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	st = strings.TrimSpace(st)
	st = st[1 : len(st)-1]
	chars := strings.Split(st, ",")

	alphabets := make(map[string]int)

	if len(st) > 0 {
		for _, char := range chars {
			char = strings.TrimSpace(char)

			if _, ok := alphabets[char]; ok {
				alphabets[char]++
			} else {
				alphabets[char] = 1
			}
		}
	}

	fmt.Println(len(alphabets))
}
