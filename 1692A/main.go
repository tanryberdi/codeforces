package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t, res int
	var a, b, c, d int
	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		res = 0
		fmt.Fscan(r, &a, &b, &c, &d)
		if a < b {
			res++
		}

		if a < c {
			res++
		}

		if a < d {
			res++
		}

		fmt.Println(res)
	}
}
