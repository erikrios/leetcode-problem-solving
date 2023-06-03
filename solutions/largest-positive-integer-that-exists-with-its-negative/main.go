package main

import "fmt"

func main() {
	fmt.Println(findMaxK([]int{-1, 2, -3, 3}))
	fmt.Println(findMaxK([]int{-1, 10, 6, 7, -7, 1}))
	fmt.Println(findMaxK([]int{-10, 8, 6, 7, -2, -3}))
}

func findMaxK(nums []int) int {
	mapperPos := make(map[int]struct{}, len(nums))
	mapperNeg := make(map[int]struct{}, len(nums))

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num > 0 {
			mapperPos[num] = struct{}{}
		} else {
			mapperNeg[num] = struct{}{}
		}
	}

	max := -1

	for k, _ := range mapperNeg {
		if _, ok := mapperPos[-k]; ok {
			if max < -k {
				max = -k
			}
		}
	}

	return max
}
