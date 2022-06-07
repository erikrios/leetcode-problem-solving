package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	prev := nums[0]

	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if prev == nums[i] {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}

		prev = nums[i]
	}

	return len(nums)
}
