package main

import "fmt"

func main() {
	var n int
	var h, g [31]int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i], &g[i])
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i != j) && (h[i] == g[j]) {
				result++
			}
		}
	}

	fmt.Println(result)
}
