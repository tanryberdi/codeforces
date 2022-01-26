package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &n)
		if n%4 == 0 {
			fmt.Fprintln(w, "YES")
			for j := 2; j <= n; j += 2 {
				fmt.Fprint(w, j, " ")
			}

			for j := 1; j < n-2; j += 2 {
				fmt.Fprint(w, j, " ")
			}
			fmt.Fprintln(w, n+n/2-1)
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
