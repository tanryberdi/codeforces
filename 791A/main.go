package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var a, b int
	fmt.Scan(&a, &b)

	result := 0
	for {
		a *= 3
		b *= 2
		result++

		if a > b {
			break
		}
	}

	fmt.Println(result)
}
