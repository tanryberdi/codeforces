package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func main() {
	var n, t int
	var a [3][]int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &t)
		a[t-1] = append(a[t-1], i+1)
	}

	result := min(len(a[0]), min(len(a[1]), len(a[2])))

	fmt.Fprintln(w, result)
	for i := 0; i < result; i++ {
		fmt.Fprintln(w, a[0][i], a[1][i], a[2][i])
	}
}
