package main

import "fmt"

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs([]int{1, 1, 1, 1}))
	fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
}

func numIdenticalPairs(nums []int) int {
	var count int

	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				count++
			}
		}
	}

	return count
}
