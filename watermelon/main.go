package main

import "fmt"

func main() {
	fmt.Println(isDividable(readInput()))
}

func isDividable(weight int) string {

	if (weight%2 == 0) && (weight > 2) {
		return "YES"
	}
	return "NO"
}

func readInput() int {
	var weight int
	fmt.Scan(&weight)
	return weight
}
