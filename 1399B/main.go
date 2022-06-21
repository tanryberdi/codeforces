package main

import (
	"bufio"
	"fmt"
	"os"
)

func findMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func calculateSteps(a []int64, b []int64) int64 {
	var steps int64
	minA := a[0]
	minB := b[0]

	for i := 1; i < len(a); i++ {
		if minA > a[i] {
			minA = a[i]
		}
	}

	for i := 1; i < len(a); i++ {
		if minB > b[i] {
			minB = b[i]
		}
	}

	for i := 0; i < len(a); i++ {
		steps += findMax(a[i]-minA, b[i]-minB)
	}

	return steps
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		a := make([]int64, n)
		b := make([]int64, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &b[i])
		}

		fmt.Fprintln(w, calculateSteps(a, b))

	}
}
