package main

import (
	"fmt"
	"strings"
)

func main() {
	readInput()

}

func readInput() {
	var st1, st2 string

	fmt.Scan(&st1, &st2)

	st1 = strings.ToLower(st1)
	st2 = strings.ToLower(st2)

	if st1 == st2 {
		fmt.Println(0)
	}

	if st1 < st2 {
		fmt.Println(-1)
	}

	if st1 > st2 {
		fmt.Println(1)
	}

}
