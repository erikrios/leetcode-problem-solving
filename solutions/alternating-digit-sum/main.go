package main

import "fmt"

func main() {
	fmt.Println(alternateDigitSum(521))
	fmt.Println(alternateDigitSum(111))
	fmt.Println(alternateDigitSum(886996))
	fmt.Println(alternateDigitSum(25))
}

func alternateDigitSum(n int) int {
	var firstSum, secondSum int
	var isPlus bool

	for n > 0 {
		digit := n % 10
		n /= 10

		if isPlus {
			firstSum += digit
			secondSum -= digit
		} else {
			firstSum -= digit
			secondSum += digit
		}

		isPlus = !isPlus
	}

	if !isPlus {
		return firstSum
	}

	return secondSum
}
