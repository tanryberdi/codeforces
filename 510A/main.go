package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &m)
	counter := 0
	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			for j := 1; j <= m; j++ {
				fmt.Fprint(w, "#")
			}
		} else {
			if counter == 0 {
				for j := 1; j <= m-1; j++ {
					fmt.Fprint(w, ".")
				}
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, "#")
				for j := 2; j <= m; j++ {
					fmt.Fprint(w, ".")
				}
			}
			counter++
			counter %= 2
		}

		fmt.Fprintln(w)
	}
}
