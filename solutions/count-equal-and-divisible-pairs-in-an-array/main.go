package main

import "fmt"

func main() {
	fmt.Println(countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2))
	fmt.Println(countPairs([]int{1, 2, 3, 4}, 1))
	fmt.Println(countPairs([]int{5, 5, 9, 2, 5, 5, 9, 2, 2, 5, 5, 6, 2, 2, 5, 2, 5, 4, 3}, 7))
}

func countPairs(nums []int, k int) int {
	var counter int

	for i := 0; i < len(nums)-1; i++ {
		fNum := nums[i]

		for j := i + 1; j < len(nums); j++ {
			sNum := nums[j]

			if fNum == sNum && i*j%k == 0 {
				counter++
			}
		}
	}

	return counter
}
