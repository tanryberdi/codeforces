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
		fmt.Fprintln(w, ((n-1)/2)+1)
	}
}
