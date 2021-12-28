package main

import "fmt"

func main() {
	var n int
	var result string

	fmt.Scan(&n)

	for i := 1; i < n; i++ {
		if i%2 == 1 {
			result += "I hate that "
		} else {
			result += "I love that "
		}
	}

	if n%2 == 1 {
		result += "I hate it"
	} else {
		result += "I love it"
	}
	fmt.Println(result)
}
