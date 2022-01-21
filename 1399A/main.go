package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t, n int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &n)

		a := make([]int, n)
		for j := range a {
			fmt.Fscan(r, &a[j])
		}
		sort.Ints(a)

		bb := true
		for j := 0; j < n-1; j++ {
			bb = bb && ((a[j+1] - a[j]) <= 1)
		}

		if bb {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}
	}
}
