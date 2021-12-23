package main

import "fmt"

func main() {
	var result, x int

	fmt.Scan(&x)

	if x%5 > 0 {
		result++
	}

	result += x / 5

	fmt.Println(result)
}
