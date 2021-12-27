package main

import "fmt"

func main() {
	var n int64
	var result int64

	fmt.Scan(&n)

	result = n / 2
	if n%2 == 1 {
		result -= n
	}

	fmt.Println(result)
}
