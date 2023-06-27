package main

import "fmt"

func main() {
	fmt.Println(maximizeSum([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(maximizeSum([]int{5, 5, 5, 5, 5}, 2))
}

func maximizeSum(nums []int, k int) int {
	var max int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > max {
			max = num
		}
	}

	var score int
	for i := 0; i < k; i++ {
		score += max
		max++
	}

	return score
}
