package main

import "fmt"

func main() {
	var n int
	var s string
	var numD, numA int

	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 0; i < n; i++ {
		if s[i] == 65 {
			numA++
		} else {
			numD++
		}
	}

	if numA > numD {
		fmt.Println("Anton")
	}

	if numA < numD {
		fmt.Println("Danik")
	}

	if numA == numD {
		fmt.Println("Friendship")
	}
}
