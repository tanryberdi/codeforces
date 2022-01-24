package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(number int) bool {

	if number < 2 {
		return false
	}

	if number == 2 {
		return true
	}

	if number%2 == 0 {
		return false
	}

	squareRoot := int(math.Sqrt(float64(number)))

	for i := 3; i <= squareRoot; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	var n int

	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n)

	for i := 4; i <= n/2; i++ {
		if !isPrime(i) && !isPrime(n-i) {
			fmt.Fprintln(w, i, n-i)
			break
		}
	}
}
