package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var n, k, q, sum int
		var a, b []int

		// Reading n and k for every test case
		fmt.Fscan(r, &n, &k)
		// Reading array A
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &q)
			a = append(a, q)
		}

		//Reading array B
		for j := 0; j < n; j++ {
			fmt.Fscan(r, &q)
			b = append(b, q)
		}

		sort.Ints(a)
		sort.Slice(b, func(i, j int) bool {
			return b[i] > b[j]
		})

		changed := 0
		for j := 0; j < n; j++ {
			if changed == k {
				break
			}

			if b[j] > a[j] {
				changed++
				a[j], b[j] = b[j], a[j]
			} else {
				break
			}
		}

		for j := 0; j < n; j++ {
			sum += a[j]
		}
		fmt.Fprintln(w, sum)
	}

}
