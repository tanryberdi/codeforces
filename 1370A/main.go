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

		if n%2 != 0 {
			n--
		}
		fmt.Fprintln(w, n/2)
	}
}
