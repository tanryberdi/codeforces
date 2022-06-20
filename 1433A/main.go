package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNumPress(n int) int {
	var steps, apart int
	apart = n
	for apart > 9 {
		apart /= 10
	}
	steps = (apart - 1) * 10
	digit := apart

	counter := 1
	for apart <= n {
		steps += counter
		apart = apart*10 + digit
		counter++
	}
	return steps
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		fmt.Fprintln(w, countNumPress(n))
	}
}
