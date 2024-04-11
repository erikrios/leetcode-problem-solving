package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber(nums []int) int {
	n := len(nums)
	numExists := make([]bool, n+1)

	for i := 0; i < n; i++ {
		numExists[nums[i]] = true
	}

	for i, exist := range numExists {
		if !exist {
			return i
		}
	}

	return -1
}
