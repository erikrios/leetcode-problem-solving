package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
}

func firstMissingPositive(nums []int) int {
	var i int
	n := len(nums)

	for i < n {
		num := nums[i]
		if num >= 1 && num <= n && nums[num-1] != nums[i] {
			nums[i], nums[num-1] = nums[num-1], nums[i]
		} else {
			i++
		}
	}

	i = 0
	for i < n && nums[i] == i+1 {
		i++
	}

	return i + 1
}
