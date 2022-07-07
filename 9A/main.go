package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	if a == b {
		return a
	}

	result := 1
	for {
		for i := 2; i <= max(a, b); i++ {
			if a%i == 0 && b%i == 0 {
				result *= i
				a /= i
				b /= i
				break
			}

			if a%i == 0 {
				a /= i
				break
			}

			if b%i == 0 {
				b /= i
				break
			}
		}

		if a == 1 && b == 1 {
			break
		}
	}

	return result
}

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var yakko, wakko int
	fmt.Fscan(r, &yakko, &wakko)

	allEvent := 6
	goodEvent := 6 - max(yakko, wakko) + 1
	fmt.Fprintf(w, "%d/%d\n", goodEvent/gcd(goodEvent, allEvent), allEvent/gcd(goodEvent, allEvent))
}
