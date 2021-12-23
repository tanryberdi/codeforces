package main

import "fmt"

func main() {
	var n, t int
	var s string

	fmt.Scan(&n, &t)
	fmt.Scan(&s)

	if len(s) > 1 {
		for i := 0; i < t; i++ {
			j := 0
			for {
				if (s[j] == 66) && (s[j+1] == 71) {
					s = s[:j] + "G" + s[j+1:]
					s = s[:j+1] + "B" + s[j+2:]
					j += 2
				} else {
					j++
				}
				if j >= len(s)-1 {
					break
				}
			}
		}
	}

	fmt.Println(s)
}
