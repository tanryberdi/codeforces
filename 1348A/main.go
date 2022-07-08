package main

import (
	"bufio"
	"fmt"
	"os"
)

func calculateNthDegreeOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return 2 * calculateNthDegreeOfTwo(n-1)
}

func absoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var t, n int
	for fmt.Fscan(r, &t); t > 0; t-- {
		fmt.Fscan(r, &n)
		nthDegreeOfTwo := calculateNthDegreeOfTwo(n)
		number1, number2 := 0, 0
		sum1, sum2 := 0, 0
		for nthDegreeOfTwo > 1 {
			if (sum1 > sum2) && (number2 < n/2) {
				sum2 += nthDegreeOfTwo
				number2++
				nthDegreeOfTwo /= 2
			} else {
				sum1 += nthDegreeOfTwo
				number1++
				nthDegreeOfTwo /= 2
			}
		}
		fmt.Fprintln(w, absoluteValue(sum1-sum2))
	}

}
