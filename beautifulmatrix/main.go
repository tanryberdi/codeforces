package main

import "fmt"

func main() {
	readInput()
}

func absoluteValue(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func readInput() {
	var a [5][5]int
	var x, y int

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&a[i][j])
			if a[i][j] == 1 {
				x = i
				y = j
			}
		}
	}

	result := 0
	result = absoluteValue(x-2) + absoluteValue(y-2)
	fmt.Println(result)
}
