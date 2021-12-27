package main

import "fmt"

func main() {
	var st1, st2 string
	var result string

	fmt.Scan(&st1)
	fmt.Scan(&st2)

	for i := 0; i < len(st1); i++ {
		if st1[i] != st2[i] {
			result += "1"
		} else {
			result += "0"
		}
	}

	fmt.Println(result)
}
