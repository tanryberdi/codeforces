package main

import "fmt"

func main() {
	readInput()

}

func readInput() {
	var n, result int
	var op string

	fmt.Scan(&n)
	result = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&op)
		if op[1] == 43 {
			result++
		} else {
			result--
		}
	}

	fmt.Println(result)
}
