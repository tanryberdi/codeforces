package main

import "fmt"

func main() {
	var k, l, m, n int
	var d, result int

	fmt.Scan(&k, &l, &m, &n)
	fmt.Scan(&d)

	result = d
	for i := 1; i <= d; i++ {
		if (i%k != 0) && (i%l != 0) && (i%m != 0) && (i%n != 0) {
			result--
		}
	}

	fmt.Println(result)
}
