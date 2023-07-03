package main

import "fmt"

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, 2}))
	fmt.Println(minStartValue([]int{1, -2, -3}))
	fmt.Println(minStartValue([]int{5, 4, -5, -5, 0}))
}

func minStartValue(nums []int) int {
	startValue := 1
	min := 101

	for i := 0; i < len(nums); i++ {
		nums[i] += startValue
		startValue = nums[i]
		if startValue < min {
			min = startValue
		}
	}

	if min > 0 {
		return 1
	}

	return 1 - min + 1
}
