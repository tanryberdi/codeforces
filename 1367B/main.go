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
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		a := make([]int, n)
		remainder := make([]int, 2)
		for i := range a {
			fmt.Fscan(r, &a[i])
			if i%2 != a[i]%2 {
				remainder[a[i]%2]++
			}
		}

		if remainder[0] == remainder[1] {
			fmt.Fprintln(w, remainder[0])
		} else {
			fmt.Fprintln(w, "-1")
		}

	}
}
