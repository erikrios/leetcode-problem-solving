package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{-2, -1}))
	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return longest(nums, 0, -1, dp)
}

func longest(nums []int, i int, prevI int, dp [][]int) int {
	n := len(nums)

	if i >= n {
		return 0
	}

	if dp[i][prevI+1] != -1 {
		return dp[i][prevI+1]
	}

	num := nums[i]

	takeLength, dontTakeLength := 0, longest(nums, i+1, prevI, dp)
	if prevI == -1 || num > nums[prevI] {
		takeLength = 1 + longest(nums, i+1, i, dp)
	}

	maxLength := takeLength
	if dontTakeLength > takeLength {
		maxLength = dontTakeLength
	}

	dp[i][prevI+1] = maxLength

	return maxLength
}
