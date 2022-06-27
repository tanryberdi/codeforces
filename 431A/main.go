package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	numbers := strings.Split(strings.TrimSpace(input1), " ")

	input2, _ := reader.ReadString('\n')
	s := strings.TrimSpace(input2)

	sum := 0
	for _, v := range s {
		i, _ := strconv.Atoi(string(v))

		a, _ := strconv.Atoi(numbers[i-1])
		sum += a
	}

	fmt.Print(sum)
}
