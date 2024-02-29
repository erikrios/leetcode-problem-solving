package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minPairSum([]int{3, 5, 2, 3}))
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}

func minPairSum(nums []int) int {
	sort.Ints(nums)

	n := len(nums)
	var max int

	for i := 0; i < n/2; i++ {
		sum := nums[i] + nums[n-i-1]
		if sum > max {
			max = sum
		}
	}

	return max
}
