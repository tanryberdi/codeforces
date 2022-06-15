package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, c int
	var mishka, chris int
	fmt.Fscan(r, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &m, &c)
		if m > c {
			mishka++
		} else if m < c {
			chris++
		}
	}

	if mishka > chris {
		fmt.Fprintln(w, "Mishka")
	} else if mishka < chris {
		fmt.Fprintln(w, "Chris")
	} else {
		fmt.Fprintln(w, "Friendship is magic!^^")
	}
}
