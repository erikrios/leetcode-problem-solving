package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates([]int{1, 1, 2}))
	fmt.Println(findDuplicates([]int{1}))
}

func findDuplicates(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		num := abs(nums[i])
		idx := num - 1
		if nums[idx] < 0 {
			results = append(results, num)
		}
		nums[idx] *= -1
	}

	return results
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}

	return num
}
