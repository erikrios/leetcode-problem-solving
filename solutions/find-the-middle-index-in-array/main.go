package main

import "fmt"

func main() {
	fmt.Println(findMiddleIndex([]int{2, 3, -1, 8, 4}))
	fmt.Println(findMiddleIndex([]int{1, -1, -4}))
	fmt.Println(findMiddleIndex([]int{2, 5}))
}

func findMiddleIndex(nums []int) int {
	var leftSum, rightSum int
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		rightSum -= num

		if leftSum == rightSum {
			return i
		}

		leftSum += num
	}

	return -1
}
