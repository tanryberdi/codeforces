package main

import "fmt"

func main() {
	var n, t, result int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		if t == 1 {
			result = 1
		}
	}

	if result == 1 {
		fmt.Println("HARD")
	} else {
		fmt.Println("EASY")
	}
}
