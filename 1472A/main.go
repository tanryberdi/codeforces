package main

import (
	"bufio"
	"fmt"
	"os"
)

func printResult(width, height int, n int64) string {
	var total int64
	total = 1
	for width%2 == 0 {
		width /= 2
		total *= 2
	}

	for height%2 == 0 {
		height /= 2
		total *= 2
	}

	if total >= n {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t, width, height int
	var n int64
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &width, &height, &n)
		fmt.Println(printResult(width, height, n))
	}

}
