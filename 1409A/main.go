package main

import (
	"bufio"
	"fmt"
	"os"
)

func getAbs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func getSteps(a, b int) int {
	var steps int
	if a == b {
		return 0
	}

	diff := getAbs(a, b)
	if diff%10 > 0 {
		steps = 1
	}

	steps += diff / 10

	return steps
}

func main() {
	var t, a, b int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &a, &b)
		fmt.Fprintln(w, getSteps(a, b))
	}
}
