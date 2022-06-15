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

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k, l, c, d, p, nl, np int
	fmt.Fscan(r, &n, &k, &l, &c, &d, &p, &nl, &np)

	var ans int
	ans = min(min((k*l)/nl, c*d), p/np)

	fmt.Fprintln(w, ans/n)
}
