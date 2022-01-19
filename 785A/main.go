package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, result int
	var shape string

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &shape)
		switch shape {
		case "Tetrahedron":
			result += 4
		case "Cube":
			result += 6
		case "Octahedron":
			result += 8
		case "Dodecahedron":
			result += 12
		default:
			result += 20
		}
	}

	fmt.Fprintln(w, result)
}
