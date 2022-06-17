package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, t int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		step := 0
		count := 0
		for step < n {
			count++
			if count%3 != 0 && count%10 != 3 {
				step++
			}
		}
		fmt.Fprintln(w, count)
	}

}
