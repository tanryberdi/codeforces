package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var x1, x2, x3 int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &x1, &x2, &x3)

	points := []int{x1, x2, x3}
	sort.Ints(points)
	fmt.Fprintln(w, (points[1]-points[0])+(points[2]-points[1]))
}
