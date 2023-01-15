package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countDigits(7))
	fmt.Println(countDigits(121))
	fmt.Println(countDigits(1248))
}

func countDigits(num int) int {
	var counter int

	numStr := strconv.Itoa(num)

	for i := 0; i < len(numStr); i++ {
		if numInt, _ := strconv.Atoi(string(numStr[i])); num%numInt == 0 {
			counter++
		}
	}

	return counter
}
