package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		numbers := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &numbers[i])
		}
		sort.Ints(numbers)
		diff := numbers[1] - numbers[0]
		for i := 2; i < n; i++ {
			if diff > numbers[i]-numbers[i-1] {
				diff = numbers[i] - numbers[i-1]
			}
		}
		fmt.Fprintln(w, diff)
	}
}
