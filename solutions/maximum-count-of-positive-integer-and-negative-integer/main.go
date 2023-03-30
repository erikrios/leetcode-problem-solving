package main

import "fmt"

func main() {
	fmt.Println(maximumCount([]int{-2, -1, -1, 1, 2, 3}))
	fmt.Println(maximumCount([]int{-3, -2, -1, 0, 0, 1, 2}))
	fmt.Println(maximumCount([]int{5, 20, 66, 1314}))
	fmt.Println(maximumCount([]int{-3, -2, -1}))
}

func maximumCount(nums []int) int {
	neg := search(nums, 0)
	pos := len(nums) - search(nums, 1)

	if neg > pos {
		return neg
	}

	return pos
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)

	for l < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
