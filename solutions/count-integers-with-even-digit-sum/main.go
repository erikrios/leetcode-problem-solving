package main

import "fmt"

func main() {
	fmt.Println(countEven(4))
	fmt.Println(countEven(30))
}

func countEven(num int) int {
	var counter int

	for i := 1; i <= num; i++ {
		if isDigitSumEven(i) {
			counter++
		}
	}

	return counter
}

func isDigitSumEven(num int) bool {
	var sum int

	for num > 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}

	return sum%2 == 0
}
