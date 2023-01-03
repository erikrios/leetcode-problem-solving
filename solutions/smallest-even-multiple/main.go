package main

import "fmt"

func main() {
	fmt.Println(smallestEvenMultiple(5))
	fmt.Println(smallestEvenMultiple(6))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}

	return n * 2
}
