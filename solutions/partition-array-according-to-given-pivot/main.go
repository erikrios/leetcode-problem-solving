package main

import "fmt"

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
	fmt.Println(pivotArray([]int{-3, 4, 3, 2}, 2))
}

func pivotArray(nums []int, pivot int) []int {
	var smaller, greater, equal int

	n := len(nums)
	for i := 0; i < n; i++ {
		if num := nums[i]; num < pivot {
			smaller++
		} else if num == pivot {
			equal++
		}
	}

	greater = smaller + equal
	equal = smaller
	smaller = 0
	res := make([]int, n, n)
	for i := 0; i < n; i++ {
		if num := nums[i]; num < pivot {
			res[smaller] = num
			smaller++
		} else if num > pivot {
			res[greater] = num
			greater++
		} else {
			res[equal] = num
			equal++
		}
	}

	return res
}
