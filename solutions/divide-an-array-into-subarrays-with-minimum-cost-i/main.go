package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumCost([]int{1, 2, 3, 12}))
	fmt.Println(minimumCost([]int{5, 4, 3}))
	fmt.Println(minimumCost([]int{10, 3, 1, 1}))
}

func minimumCost(nums []int) int {
	sum := nums[0]
	num1 := math.MaxInt
	num2 := math.MaxInt

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num < num1 {
			num2 = num1
			num1 = num
		} else if num < num2 {
			num2 = num
		}
	}

	return sum + num1 + num2
}
