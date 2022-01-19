package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, result, min, max, answer int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)
	fmt.Fscan(r, &result)
	min = result
	max = result
	for i := 0; i < n-1; i++ {
		fmt.Fscan(r, &result)
		if result < min {
			min = result
			answer++
		}

		if result > max {
			max = result
			answer++
		}
	}

	fmt.Fprintln(w, answer)
}
