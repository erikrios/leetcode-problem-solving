package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}

func subsets(nums []int) [][]int {
	return backtracking(nums, 0, []int{})
}

func backtracking(nums []int, i int, res []int) [][]int {
	if i == len(nums) {
		resCopy := make([]int, len(res))
		copy(resCopy, res)
		return [][]int{resCopy}
	}

	results := make([][]int, 0)
	// Not take
	result := backtracking(nums, i+1, res)
	results = append(results, result...)

	// Take
	res = append(res, nums[i])
	result = backtracking(nums, i+1, res)
	results = append(results, result...)

	return results
}
