package main

import "fmt"

func main() {
	var k, n, w int
	var result int

	fmt.Scan(&k, &n, &w)
	result = k * ((w * (w + 1)) / 2)
	result -= n
	if result < 0 {
		result = 0
	}
	fmt.Println(result)
}
