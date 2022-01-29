package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, k, y int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &k)

	counter := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &y)
		if y+k < 6 {
			counter++
		}
	}

	fmt.Fprintln(w, counter/3)
}
