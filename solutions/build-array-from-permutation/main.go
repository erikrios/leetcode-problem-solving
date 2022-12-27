package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
	fmt.Println(buildArray([]int{5, 0, 1, 2, 3, 4}))
}

func buildArray(nums []int) []int {
	results := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		results[i] = nums[nums[i]]
	}

	return results
}
