package main

import "fmt"

func main() {
	var s [4]int64

	shoes := make(map[int64]int)
	need := 4
	for i := 0; i < 4; i++ {
		fmt.Scan(&s[i])
		if _, ok := shoes[s[i]]; !ok {
			need--
			shoes[s[i]] = 1
		} else {
			shoes[s[i]]++
		}
	}

	fmt.Println(need)
}
