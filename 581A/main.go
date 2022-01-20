package main

import (
	"bufio"
	"fmt"
	"os"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	var a, b int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &a, &b)

	fmt.Fprintln(w, getMin(a, b),
		(getMax(a, b)-getMin(a, b))/2,
	)
}
