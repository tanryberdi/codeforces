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
	//fmt.Fprintln(w, t)
	for i := 0; i < t; i++ {
		var x, y, n int
		fmt.Fscan(r, &x, &y, &n)

		temp := (n / x) * x
		if temp+y <= n {
			fmt.Fprintln(w, temp+y)
		} else {
			fmt.Fprintln(w, temp+y-x)
		}

	}
}
