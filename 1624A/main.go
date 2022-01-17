package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, t int

	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &t)
	for ; t > 0; t-- {
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(reader, &a[i])
		}
		sort.Ints(a)
		fmt.Fprintln(writer, a[n-1]-a[0])
	}
}
