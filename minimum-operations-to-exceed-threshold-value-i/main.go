package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{2, 11, 10, 1, 3}, 10))
	fmt.Println(minOperations([]int{1, 1, 2, 4, 9}, 1))
	fmt.Println(minOperations([]int{1, 1, 2, 4, 9}, 9))
}

func minOperations(nums []int, k int) int {
	var counter int

	for i := 0; i < len(nums); i++ {
		if num := nums[i]; num < k {
			counter++
		}
	}

	return counter
}
