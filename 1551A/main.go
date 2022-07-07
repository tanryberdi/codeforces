package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t, n int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		if n%3 == 0 {
			fmt.Fprintln(w, n/3, n/3)
		} else {
			switch n % 3 {
			case 1:
				fmt.Fprintln(w, n/3+1, n/3)
			default:
				fmt.Fprintln(w, n/3, n/3+1)
			}
		}
	}
}
