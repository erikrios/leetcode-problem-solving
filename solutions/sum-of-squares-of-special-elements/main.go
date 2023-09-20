package main

import "fmt"

func main() {
	fmt.Println(sumOfSquares([]int{1, 2, 3, 4}))
	fmt.Println(sumOfSquares([]int{2, 7, 1, 19, 18, 3}))
}

func sumOfSquares(nums []int) int {
	n := len(nums)
	var sum int

	for i := 0; i < n; i++ {
		num := nums[i]
		if n%(i+1) == 0 {
			num *= num
			sum += num
		}
	}

	return sum
}
