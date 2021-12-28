package main

import "fmt"

func main() {
	var n int
	var p, result float64
	fmt.Scan(&n)

	result = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&p)
		result += p
	}

	fmt.Printf("%.12f", result/float64(n))
}
