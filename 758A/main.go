package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var a [100]int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)
	max := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		if max < a[i] {
			max = a[i]
		}
	}

	needSpend := 0
	for i := 0; i < n; i++ {
		needSpend += max - a[i]
	}

	fmt.Fprintln(w, needSpend)
}
