package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	memo := make([]int, n)
	for i := 0; i < n; i++ {
		memo[i] = -1
	}

	maxOne := max(nums, 0, memo)
	maxTwo := max(nums, 1, memo)

	if maxOne > maxTwo {
		return maxOne
	}

	return maxTwo
}

func max(nums []int, startIdx int, memo []int) int {
	n := len(nums)

	if startIdx >= n {
		return 0
	}

	if startIdx == n-1 || startIdx == n-2 {
		return nums[startIdx]
	}

	if memo[startIdx] != -1 {
		return memo[startIdx]
	}

	nextTwoMax := max(nums, startIdx+2, memo)
	nextThreeMax := max(nums, startIdx+3, memo)

	num := nums[startIdx]

	sum := num

	if nextTwoMax > nextThreeMax {
		sum += nextTwoMax
	} else {
		sum += nextThreeMax
	}

	memo[startIdx] = sum

	return sum
}
