package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	diff := math.MaxInt
	closest := 0

	for i := 0; i < len(nums); i++ {
		first := nums[i]
		start := i + 1
		end := len(nums) - 1

		for start < end {
			if first+nums[start]+nums[end] == target {
				return target
			} else if int(math.Abs(float64(first+nums[start]+nums[end]-target))) < diff {
				diff = int(math.Abs(float64(first + nums[start] + nums[end] - target)))
				closest = first + nums[start] + nums[end]
			}

			if first+nums[start]+nums[end] > target {
				end--
			} else {
				start++
			}
		}
	}

	return closest
}
