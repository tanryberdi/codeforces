package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSecret(st string) string {
	var secSt string

	if len(st) < 3 {
		return st
	}

	secSt += string(st[0])

	for i := 1; i < len(st)-1; i += 2 {
		secSt += string(st[i])
	}

	secSt += string(st[len(st)-1])
	return secSt
}

func main() {
	var t int
	var st string
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(r, &st)
		fmt.Fprintln(w, findSecret(st))
	}
}
