package main

import (
	"fmt"
)

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
	fmt.Println(differenceOfSum([]int{1, 2, 3, 4}))
}

func differenceOfSum(nums []int) int {
	var elementSum, digitSum int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		elementSum += num

		for num > 0 {
			digitSum += num % 10
			num /= 10
		}

	}

	var result int
	if elementSum > digitSum {
		result = elementSum - digitSum
	} else {
		result = digitSum - elementSum
	}

	return result
}
