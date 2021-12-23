package main

import "fmt"

func main() {
	var n int64
	var numLucky int

	fmt.Scan(&n)

	for n > 0 {
		if (n%10 == 4) || (n%10 == 7) {
			numLucky++
		}
		n /= 10
	}

	if (numLucky == 7) || (numLucky == 4) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
