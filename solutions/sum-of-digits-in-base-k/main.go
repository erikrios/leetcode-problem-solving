package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumBase(34, 6))
	fmt.Println(sumBase(10, 10))
	fmt.Println(sumBase(42, 2))
	fmt.Println(sumBase(25, 6))
}

func sumBase(n int, k int) int {
	var sum int

	for n > 0 {
		sum += n % k
		n /= k
	}

	return sum
}
