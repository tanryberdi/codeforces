package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var n, result int
	var st string

	fmt.Scan(&n)
	fmt.Scan(&st)

	result = 0
	for i := 0; i < len(st)-1; i++ {
		if st[i] == st[i+1] {
			result++
		}
	}

	fmt.Println(result)
}
