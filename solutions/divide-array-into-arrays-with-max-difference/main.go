package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
	fmt.Println(divideArray([]int{1, 3, 3, 2, 7, 3}, 3))
	fmt.Println(divideArray([]int{15, 13, 12, 13, 12, 14, 12, 2, 3, 13, 12, 14, 14, 13, 5, 12, 12, 2, 13, 2, 2}, 2))
}

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	results := make([][]int, 0)

	for i := 0; i < n; i += 3 {
		first := nums[i]
		second := nums[i+1]
		third := nums[i+2]

		if third-first <= k {
			results = append(results, []int{first, second, third})
		} else {
			return [][]int{}
		}
	}

	return results
}
