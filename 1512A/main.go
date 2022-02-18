package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(r, &n)
		numbers := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(r, &numbers[j])
		}

		for j := range numbers {
			if numbers[j] != numbers[(j+1)%n] && numbers[j] != numbers[(j+2)%n] {
				fmt.Fprintln(w, j+1)
				break
			}
		}
	}
}
