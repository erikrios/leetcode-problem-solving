package main

import "fmt"

func main() {
	fmt.Println(commonFactors(12, 6))
	fmt.Println(commonFactors(25, 30))
}

func commonFactors(a int, b int) int {
	var counter int

	var min int
	if a < b {
		min = a
	} else {
		min = b
	}

	for i := min; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			counter++
		}
	}

	return counter
}
