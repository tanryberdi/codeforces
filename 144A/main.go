package main

import "fmt"

func main() {
	var n, a int
	max := 0
	imax := 0
	min := 101
	imin := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a <= min {
			min = a
			imin = i
		}

		if a > max {
			max = a
			imax = i
		}
	}

	//fmt.Println("max ==> ", max, " ", imax)
	//fmt.Println("min ==> ", min, " ", imin)

	result := 0
	if imax < imin {
		result = imax + (n - imin - 1)
	} else {
		result = imax + (n - imin - 2)
	}

	fmt.Println(result)

}
