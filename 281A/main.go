package main

import (
	"fmt"
	"strings"
)

func main() {
	readInput()
}

func readInput() {
	var st string

	fmt.Scan(&st)

	st = strings.Title(st)
	fmt.Println(st)
}
