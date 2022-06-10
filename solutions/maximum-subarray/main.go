package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	maxSoFar, maxEndingHere := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = int(math.Max(float64(maxEndingHere)+float64(nums[i]), float64(nums[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}

	return maxSoFar
}
