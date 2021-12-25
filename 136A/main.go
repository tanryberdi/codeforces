package main

import "fmt"

func main() {
	var n, t int
	var a [101]int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		a[t] = i + 1
	}

	for i := 0; i < n; i++ {
		fmt.Print(a[i+1], " ")
	}
}
