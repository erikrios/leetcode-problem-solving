package main

import "fmt"

func main() {
	fmt.Println(returnToBoundaryCount([]int{2, 3, -5}))
	fmt.Println(returnToBoundaryCount([]int{3, 2, -3, -4}))
}

func returnToBoundaryCount(nums []int) int {
	var pos int
	var counter int

	for i := 0; i < len(nums); i++ {
		pos += nums[i]
		if pos == 0 {
			counter++
		}
	}

	return counter
}
