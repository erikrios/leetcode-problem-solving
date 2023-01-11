package main

import "fmt"

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}))
}

func createTargetArray(nums []int, index []int) []int {
	results := make([]int, len(nums), len(nums))

	for i, num := range nums {
		if i == index[i] {
			results[i] = num
		} else {
			insertAt(results, index[i], num)
		}
	}

	return results
}

func insertAt(nums []int, index, val int) {
	for i := len(nums) - 1 - 1; i >= index; i-- {
		nums[i+1] = nums[i]
	}

	nums[index] = val
}
