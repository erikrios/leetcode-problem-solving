package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4}))
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		return [][]int{newNums}
	}

	newResults := make([][]int, 0)
	for i, num := range nums {
		beforeNums := make([]int, len(nums[:i]))
		copy(beforeNums, nums[:i])
		selecteds := beforeNums
		afterNums := make([]int, len(nums[i+1:]))
		copy(afterNums, nums[i+1:])
		selecteds = append(selecteds, afterNums...)
		results := permute(selecteds)

		for _, result := range results {
			res := []int{num}
			res = append(res, result...)
			newResults = append(newResults, res)
		}
	}

	return newResults
}
