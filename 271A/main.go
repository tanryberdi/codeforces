package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)

	for year := y + 1; year <= 10000; year++ {
		digits := make(map[int]int)
		current := year

		for current > 0 {
			digit := current % 10
			if _, ok := digits[digit]; ok {
				digits[digit]++
			} else {
				digits[digit] = 1
			}
			current /= 10
		}

		if len(digits) == 4 {
			fmt.Println(year)
			break
		}
	}

}
