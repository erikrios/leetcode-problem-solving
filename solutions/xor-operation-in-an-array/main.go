package main

import "fmt"

func main() {
	fmt.Println(xorOperation(5, 0))
	fmt.Println(xorOperation(4, 3))
}

func xorOperation(n int, start int) int {
	var result int

	for i := 0; i < n; i++ {
		val := start + 2*i
		result ^= val
	}

	return result
}
