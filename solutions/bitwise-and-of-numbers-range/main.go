package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 0))
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}

func rangeBitwiseAnd(left int, right int) int {
	res := left

	for i := left + 1; i <= right; i++ {
		res = res & i
		if res == 0 {
			break
		}
	}

	return res
}
