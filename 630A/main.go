package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int64
	fmt.Fscan(r, &n)
	fmt.Fprintln(w, degreeOfFive(n))
}

func degreeOfFive(n int64) int {
	switch n {
	case 0:
		return 1
	case 1:
		return 5
	default:
		return 25
	}
}
