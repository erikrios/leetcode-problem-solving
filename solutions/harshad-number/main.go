package main

import "fmt"

func main() {
	fmt.Println(sumOfTheDigitsOfHarshadNumber(18))
	fmt.Println(sumOfTheDigitsOfHarshadNumber(23))
}

func sumOfTheDigitsOfHarshadNumber(x int) int {
	var digitSum int
	n := x

	for n > 0 {
		digit := n % 10
		n /= 10
		digitSum += digit
	}

	if x%digitSum == 0 {
		return digitSum
	}

	return -1
}
