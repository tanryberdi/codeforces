package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	for fmt.Fscan(r, &n); n > 0; n-- {
		var a, b int
		fmt.Fscan(r, &a, &b)
		big := max(a, b)
		small := min(a, b)
		if big == small {
			big *= 2
			fmt.Fprintln(w, big*big)
		} else {
			if big >= small*2 {
				fmt.Fprintln(w, big*big)
			} else {
				fmt.Fprintln(w, 4*small*small)
			}
		}
	}
}
