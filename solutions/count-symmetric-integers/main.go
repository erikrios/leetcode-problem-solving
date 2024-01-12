package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSymmetricIntegers(1, 100))
	fmt.Println(countSymmetricIntegers(1200, 1230))
}

func countSymmetricIntegers(low, high int) int {
	var counter int

	for i := low; i <= high; i++ {
		if digits := getDigits(i); isSymmetric(digits) {
			counter++
		}
	}

	return counter
}

func getDigits(num int) (digits []int) {
	for num > 0 {
		digit := num % 10
		num /= 10
		digits = append(digits, digit)
	}

	return
}

func isSymmetric(digits []int) bool {
	n := len(digits)
	if n%2 != 0 {
		return false
	}

	var firstSum, lastSum int
	for i := 0; i < n/2; i++ {
		firstNum := digits[i]
		lastNum := digits[n-1-i]

		firstSum += firstNum
		lastSum += lastNum
	}

	return firstSum == lastSum
}
