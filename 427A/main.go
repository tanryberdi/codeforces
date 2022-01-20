package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, h int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)
	hired := 0
	crime := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &h)
		if h >= 0 {
			hired += h
		} else {
			if hired > 0 {
				hired--
			} else {
				crime++
			}
		}
	}

	fmt.Fprintln(w, crime)
}
