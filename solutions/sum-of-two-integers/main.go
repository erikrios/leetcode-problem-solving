package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(2, 3))
}

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b) << 1

		a = sum
		b = carry
	}

	return a
}
