package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t, candy int
	for fmt.Fscan(r, &t); t > 0; t-- {
		n1 := 0
		n2 := 0
		for fmt.Fscan(r, &n); n > 0; n-- {
			fmt.Fscan(r, &candy)
			switch candy {
			case 1:
				n1++
			default:
				n2++
			}
		}

		if n1 > 0 {
			if n1%2 == 0 {
				fmt.Fprintln(w, "YES")
			} else {
				fmt.Fprintln(w, "NO")
			}
		} else {
			if n2%2 == 0 {
				fmt.Fprintln(w, "YES")
			} else {
				fmt.Fprintln(w, "NO")
			}
		}
	}
}
