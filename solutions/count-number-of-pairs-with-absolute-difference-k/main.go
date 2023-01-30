package main

import "fmt"

func main() {
	fmt.Println(countDifference([]int{1, 2, 2, 1}, 1))
	fmt.Println(countDifference([]int{1, 3}, 3))
	fmt.Println(countDifference([]int{3, 2, 1, 5, 4}, 2))
}

func countDifference(nums []int, k int) int {
	var counter int

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			val := nums[i] - nums[j]
			if val < 0 {
				val *= -1
			}

			if val == k {
				counter++
			}
		}
	}

	return counter
}
