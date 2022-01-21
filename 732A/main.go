package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, b int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &k, &b)
	buyed := 1
	for {
		price := buyed * k
		if (price%10 == 0) || (price%10 == b) {
			break
		}
		buyed++
	}

	fmt.Fprintln(w, buyed)
}
