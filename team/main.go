package main

import "fmt"

func main() {
	fmt.Println(readInput())
}

func readInput() int {
	var nProblems int
	var petya, vasya, tonya int

	fmt.Scan(&nProblems)

	nSure := 0
	for i := 0; i < nProblems; i++ {
		fmt.Scan(&petya, &vasya, &tonya)
		if petya+vasya+tonya > 1 {
			nSure++
		}
	}

	return nSure
}
