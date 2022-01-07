package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var st string
	fmt.Scan(&n, &st)

	st = strings.ToLower(st)

	alphabets := make(map[string]int)
	for i := 0; i < n; i++ {
		if _, ok := alphabets[string(st[i])]; !ok {
			alphabets[string(st[i])] = 1
		} else {
			alphabets[string(st[i])]++
		}
	}

	if len(alphabets) >= 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
