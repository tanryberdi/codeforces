package main

import "fmt"

func main() {
	var n, h, height int
	var result int

	fmt.Scan(&n, &h)
	for i := 0; i < n; i++ {
		fmt.Scan(&height)

		if height > h {
			result += 2
		} else {
			result += 1
		}
	}

	fmt.Println(result)
}
