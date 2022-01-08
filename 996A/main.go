package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	result := n / 100
	n = n % 100

	result += n / 20
	n %= 20

	result += n / 10
	n %= 10

	result += n / 5
	n %= 5

	result += n

	fmt.Println(result)
}
