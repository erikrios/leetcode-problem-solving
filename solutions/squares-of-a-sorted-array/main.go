package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n, n)
	i, j := 0, n-1
	for k := n - 1; k >= 0; k-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[k] = nums[i] * nums[i]
			i++
		} else {
			res[k] = nums[j] * nums[j]
			j--
		}
	}

	return res
}
