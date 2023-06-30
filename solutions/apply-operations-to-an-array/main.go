package main

import "fmt"

func main() {
	fmt.Println(applyOperations([]int{1, 2, 2, 1, 1, 0}))
	fmt.Println(applyOperations([]int{0, 1}))
}

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		num := nums[i]
		next := nums[i+1]

		if num == next {
			nums[i] = num * 2
			nums[i+1] = 0
		}
	}

	res := make([]int, len(nums), len(nums))
	firstPtr := 0
	secondPtr := len(res) - 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num == 0 {
			res[secondPtr] = num
			secondPtr--
		} else {
			res[firstPtr] = num
			firstPtr++
		}
	}

	return res
}
