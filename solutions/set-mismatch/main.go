package main

import (
	"fmt"
)

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
	fmt.Println(findErrorNums([]int{1, 1}))
	fmt.Println(findErrorNums([]int{1, 2, 5, 4, 5}))
}

func findErrorNums(nums []int) []int {
	n := len(nums)

	sets := make([]int, n)

	res := []int{-1, -1}
	for i := 0; i < n; i++ {
		num := nums[i]
		sets[num-1]++
	}

	for i := 0; i < n; i++ {
		if set := sets[i]; set == 0 {
			res[1] = i + 1
		} else if set == 2 {
			res[0] = i + 1
		}

		if res[0] != -1 && res[1] != -1 {
			break
		}
	}

	return res
}
