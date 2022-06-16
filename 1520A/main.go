package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t int
	var line string
	for fmt.Fscan(r, &t); t > 0; t-- {
		var char rune
		isOk := true
		fmt.Fscan(r, &n)
		fmt.Fscan(r, &line)
		chars := make(map[rune]int)

		for _, c := range line {
			char = c
			chars[c]++
			break
		}

		for _, c := range line {
			if char == c {
				chars[c]++
			} else {
				if _, ok := chars[c]; ok {
					isOk = false
					break
				} else {
					char = c
					chars[c]++
				}
			}
		}

		if isOk {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
