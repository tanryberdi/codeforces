package main

import "fmt"

func main() {
	var st1, st2 string

	fmt.Scan(&st1, &st2)

	if len(st1) != len(st2) {
		fmt.Println("NO")
	} else {
		isTrue := true
		for i := 0; i < len(st1); i++ {
			isTrue = isTrue && (st1[i] == st2[len(st2)-1-i])
			//fmt.Println(i, "-->", len(st2)-1-i)
		}

		if isTrue {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
