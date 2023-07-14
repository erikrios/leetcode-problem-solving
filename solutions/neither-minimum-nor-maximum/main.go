package main

import "fmt"

func main() {
	fmt.Println(findNonMinOrMax([]int{3, 2, 1, 4}))
	fmt.Println(findNonMinOrMax([]int{1, 2}))
	fmt.Println(findNonMinOrMax([]int{2, 1, 3}))
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	min, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if num < min {
			min, num = num, min
		}

		if num > max {
			max, num = num, max
		}

		if num > min && num < max {
			return num
		}
	}

	return -1
}
