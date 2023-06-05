package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{1, 1, 1}))
	fmt.Println(minOperations([]int{1, 5, 2, 4, 1}))
	fmt.Println(minOperations([]int{8}))
}

func minOperations(nums []int) int {
	var counter int

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		num := nums[i]

		if num <= prev {
			target := prev + 1
			counter += target - num
			nums[i] = target
		}
	}

	return counter
}
