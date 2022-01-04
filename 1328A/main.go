package main

import "fmt"

func main() {
	var n int
	var a, b int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		if a%b == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(b - a%b)
		}
	}
}
