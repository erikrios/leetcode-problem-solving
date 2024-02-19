package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minOperations([]int{2, 3, 3, 2, 2, 4, 2, 3, 4}))
	fmt.Println(minOperations([]int{2, 1, 2, 2, 3, 3}))
	fmt.Println(minOperations([]int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}))
}

func minOperations(nums []int) int {
	numTracker := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numTracker[nums[i]]++
	}

	var operations int
	for _, v := range numTracker {
		if v == 1 {
			return -1
		}
		operations += int(math.Ceil(float64(v) / 3))
	}

	return operations
}
