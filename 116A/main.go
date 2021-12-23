package main

import "fmt"

func main() {
	var n int
	var a, b, inBus, max int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		inBus -= a
		inBus += b
		if inBus > max {
			max = inBus
		}
	}

	fmt.Println(max)
}
