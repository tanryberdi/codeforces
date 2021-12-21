package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var n, m int
	fmt.Scan(&m, &n)
	result := (n * m) / 2
	fmt.Println(result)
}
