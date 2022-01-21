package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var x []int
	var x1, x2, x3, x4 int
	var a, b, c int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &x1, &x2, &x3, &x4)
	x = append(x, x1)
	x = append(x, x2)
	x = append(x, x3)
	x = append(x, x4)
	sort.Ints(x)

	c = x[3] - x[0]
	b = x[1] - c
	a = x[0] - b
	fmt.Fprintln(w, a, b, c)
}
