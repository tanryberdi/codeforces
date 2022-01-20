package main

import (
	"bufio"
	"fmt"
	"os"
)

func getNumOfTask(n, k int) int {
	reamainingTime := 240 - k
	if reamainingTime <= 0 {
		return 0
	}

	sum := 0
	i := 0
	answer := 0
	for i = 1; i <= n; i++ {
		if sum+i*5 <= reamainingTime {
			sum += i * 5
			answer = i
		} else {
			break
		}
	}

	return answer
}

func main() {
	var n, k int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &k)
	fmt.Fprintln(w, getNumOfTask(n, k))
}
