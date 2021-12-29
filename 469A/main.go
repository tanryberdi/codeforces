package main

import "fmt"

func main() {
	var n, p, q, r int
	var round [101]int

	fmt.Scan(&n)
	fmt.Scan(&p)
	for i := 0; i < p; i++ {
		fmt.Scan(&r)
		round[r] = 1
	}

	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&r)
		round[r] = 1
	}

	bb := true
	for i := 1; i <= n; i++ {
		bb = bb && (round[i] == 1)
	}

	if bb {
		fmt.Println("I become the guy.")
	} else {
		fmt.Println("Oh, my keyboard!")
	}
}
