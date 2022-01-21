package main

import (
	"bufio"
	"fmt"
	"os"
)

func tenPower(pow int) int {
	var result int
	if pow == 0 {
		return 1
	}

	result = 1
	for i := 0; i < pow; i++ {
		result *= 10
	}

	return result
}

func main() {
	var t, n int
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &t)
	for i := 0; i < t; i++ {
		var numbers []int
		fmt.Fscan(r, &n)

		digits := 0
		newNumber := n

		for {
			if newNumber%10 > 0 {
				if digits > 0 {
					lastNum := (newNumber % 10) * tenPower(digits)
					//fmt.Fprintln(w, "lastNum=", lastNum)
					numbers = append(numbers, lastNum)
					n -= lastNum
					newNumber = n
					digits = 0
				} else {
					lastNum := newNumber % 10
					//fmt.Fprintln(w, "lastNum=", lastNum)
					numbers = append(numbers, lastNum)
					n -= lastNum
					newNumber = n
				}
			} else {
				newNumber /= 10
				digits++
			}

			if n == 0 {
				break
			}
		}
		fmt.Fprintln(w, len(numbers))
		for j := 0; j < len(numbers)-1; j++ {
			fmt.Fprint(w, numbers[j], " ")
		}
		fmt.Fprintln(w, numbers[len(numbers)-1])
		numbers = nil

	}
}
