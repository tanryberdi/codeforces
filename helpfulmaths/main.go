package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	numbers := readInput()
	result := prepExp(numbers)
	fmt.Println(result)
}

func prepExp(numbers []string) string {
	exp := ""
	if len(numbers) < 2 {
		return numbers[0]
	}

	for i := 0; i < len(numbers); i++ {
		exp += numbers[i]
		if i != len(numbers)-1 {
			exp += "+"
		}
	}

	return exp
}

func readInput() []string {
	var st string
	fmt.Scan(&st)

	numbers := strings.Split(st, "+")
	sort.Strings(numbers)

	return numbers
}
