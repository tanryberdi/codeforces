package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int

	fmt.Fscan(r, &n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	sereja := 0
	dima := 0
	head := 0
	tail := len(a) - 1
	for {
		// Sereja's turn
		if a[head] > a[tail] {
			sereja += a[head]
			head++
			if head > tail {
				break
			}
		} else {
			sereja += a[tail]
			tail--
			if head > tail {
				break
			}
		}

		// Dima's turn
		if a[head] > a[tail] {
			dima += a[head]
			head++
			if head > tail {
				break
			}
		} else {
			dima += a[tail]
			tail--
			if head > tail {
				break
			}
		}
	}

	fmt.Fprintln(w, sereja, dima)
}
