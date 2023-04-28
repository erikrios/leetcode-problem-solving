package main

import "fmt"

func main() {
	fmt.Println(sumOfMultiples(7))
	fmt.Println(sumOfMultiples(10))
	fmt.Println(sumOfMultiples(9))
}

func sumOfMultiples(n int) int {
	if n == 1 {
		return 0
	}

	if n%7 == 0 || n%5 == 0 || n%3 == 0 {
		return sumOfMultiples(n-1) + n
	}

	return sumOfMultiples(n - 1)
}
