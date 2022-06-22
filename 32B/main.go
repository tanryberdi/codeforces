package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var st, output string
	fmt.Fscan(r, &st)
	index := 0

	for {
		switch st[index] {
		case '.':
			output += "0"
			index++
		case '-':
			if st[index+1] == '.' {
				output += "1"
				index += 2
			} else {
				output += "2"
				index += 2
			}
		}

		if index >= len(st) {
			break
		}
	}
	fmt.Fprintln(w, output)
}
