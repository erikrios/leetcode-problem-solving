package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1, 2, 3}, 0))
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	var counter int

	for i := 0; i < len(nums); i++ {
		product := 1
		for j := i + 1; j <= len(nums); j++ {
			subArrays := nums[i:j]
			if product *= subArrays[len(subArrays)-1]; product < k {
				counter++
			} else {
				break
			}
		}
	}

	return counter
}
