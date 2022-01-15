package main

import (
	"fmt"
)

func main() {
	var st1, st2, st3 string

	fmt.Scan(&st1, &st2, &st3)

	alphabets1 := make(map[string]int)
	for i := 0; i < len(st1); i++ {
		alphabets1[string(st1[i])]++
	}

	for i := 0; i < len(st2); i++ {
		alphabets1[string(st2[i])]++
	}

	alphabets2 := make(map[string]int)
	for i := 0; i < len(st3); i++ {
		alphabets2[string(st3[i])]++
	}

	isEqual := true
	for key := range alphabets1 {
		isEqual = isEqual && (alphabets1[key] == alphabets2[key])
	}

	for key := range alphabets2 {
		isEqual = isEqual && (alphabets1[key] == alphabets2[key])
	}

	if isEqual {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
