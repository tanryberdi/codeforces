package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var up, low int
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		if (s[i] > 64) && (s[i] < 91) {
			up++
		} else {
			low++
		}
	}

	if low >= up {
		s = strings.ToLower(s)
	} else {
		s = strings.ToUpper(s)
	}

	fmt.Println(s)
}
