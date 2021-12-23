package main

import "fmt"

func main() {
	readInput()
}

func readInput() {
	var st string

	fmt.Scan(&st)

	chars := make(map[int]int)
	for i := 0; i < len(st); i++ {
		if _, ok := chars[int(st[i])]; ok {
			chars[int(st[i])]++
		} else {
			chars[int(st[i])] = 1
		}
	}

	if len(chars)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}
