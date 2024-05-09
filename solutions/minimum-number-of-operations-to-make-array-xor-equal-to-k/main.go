package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperations([]int{2, 0, 2, 0}, 0))
}

func minOperations(nums []int, k int) int {
	xorTotal := nums[0]

	for i := 1; i < len(nums); i++ {
		xorTotal ^= nums[i]
	}

	if xorTotal == k {
		return 0
	}

	diff := xorTotal ^ k
	var count int

	for diff > 0 {
		count++
		diff = diff & (diff - 1)
	}

	return count
}
