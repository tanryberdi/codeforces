package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var n, k, a int
	fmt.Scan(&n, &k)

	fmt.Scan(&a)
	border := a
	numPart := 0
	if border > 0 {
		numPart = 1
	} else {
		numPart = 0
	}

	//fmt.Println(a, " ", border, " ", numPart)

	for i := 1; i < n; i++ {
		fmt.Scan(&a)
		if i <= k-1 {
			border = a
		}

		if (a >= border) && (a > 0) {
			numPart++
		}
		//fmt.Println("i=", i, " a=", a, " border=", border, " numpart=", numPart)
	}

	fmt.Println(numPart)
}
