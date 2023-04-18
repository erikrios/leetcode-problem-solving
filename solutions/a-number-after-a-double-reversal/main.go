package main

import "fmt"

func main() {
	fmt.Println(isSameAfterReversals(526))
	fmt.Println(isSameAfterReversals(1800))
	fmt.Println(isSameAfterReversals(0))
}

func isSameAfterReversals(num int) bool {

	reversed := reverse(num)
	reversed = reverse(reversed)

	return num == reversed
}

func reverse(num int) int {
	var reversed int
	tempNum := num

	digits := make([]int, 0)
	multiplicator := 0.1
	for tempNum > 0 {
		digit := tempNum % 10
		tempNum /= 10
		digits = append(digits, digit)
		multiplicator *= 10
	}

	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		reversed += digit * int(multiplicator)
		multiplicator /= 10
	}

	return reversed
}
