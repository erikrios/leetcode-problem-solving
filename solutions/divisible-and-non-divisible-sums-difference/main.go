package main

import "fmt"

func main() {
	fmt.Println(differenceOfSums(10, 3))
	fmt.Println(differenceOfSums(5, 6))
	fmt.Println(differenceOfSums(5, 1))
}

// Iterative ways
func differenceOfSums(n, m int) int {
	var result int

	for i := 1; i <= n; i++ {
		if i%m == 0 {
			result -= i
		} else {
			result += i
		}
	}

	return result
}

// Recursive ways
func differenceOfSum(n, m int) int {
	if n == 0 {
		return 0
	}

	var result int

	if n%m == 0 {
		result = -n
	} else {
		result = n
	}

	return result + differenceOfSum(n-1, m)
}
