package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
	fmt.Println(countPairs([]int{-6, 2, 5, -2, -7, -1, 3}, -2))
}

func countPairs(nums []int, target int) int {
	var total int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				total++
			}
		}
	}

	return total
}
