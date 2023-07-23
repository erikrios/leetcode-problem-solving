package main

import "fmt"

func main() {
	fmt.Println(findKDistantIndices([]int{3, 4, 9, 1, 3, 9, 5}, 9, 1))
	fmt.Println(findKDistantIndices([]int{2, 2, 2, 2, 2}, 2, 2))
	fmt.Println(findKDistantIndices([]int{1, 1000, 1, 1000}, 1, 1))
	fmt.Println(findKDistantIndices([]int{1, 1, 2, 2, 2}, 2, 1))
}

func findKDistantIndices(nums []int, key, k int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if num := nums[i]; num == key {
			startIdx := i - k
			endIdx := i + k

			if startIdx < 0 {
				startIdx = 0
			}

			if len(res) > 0 && res[len(res)-1] >= startIdx {
				startIdx = res[len(res)-1] + 1
			}

			if endIdx > len(nums)-1 {
				endIdx = len(nums) - 1
			}

			for j := startIdx; j <= endIdx; j++ {
				res = append(res, j)
			}
		}
	}

	return res
}
