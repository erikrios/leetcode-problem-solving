package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumSum(2932))
	fmt.Println(minimumSum(4009))
}

func minimumSum(num int) int {
	var digits [4]int
	var cur int

	for num > 0 {
		digits[cur] = num % 10
		cur++
		num /= 10
	}

	digits = bubbleSort(digits)

	num1 := digits[0]*10 + digits[2]
	num2 := digits[1]*10 + digits[3]

	return num1 + num2
}

func bubbleSort(digits [4]int) [4]int {
	for i := 0; i < len(digits)-1; i++ {
		for j := 0; j < len(digits)-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}

	return digits
}
