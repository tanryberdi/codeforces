package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func addTwos(answer []string, times int) []string {
	for i := 0; i < times; i++ {
		answer = append(answer, "2")
	}

	return answer
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var num int
	fmt.Fscan(r, &num)
	var ans []string

	twos := num / 2
	divTwo := num%2 == 0

	if divTwo {
		ans = addTwos(ans, twos)
	} else {
		ans = addTwos(ans, twos-1)
		ans = append(ans, "3")
	}

	fmt.Fprintln(w, len(ans))
	fmt.Fprintf(w, "%v", strings.Join(ans, " "))
}
