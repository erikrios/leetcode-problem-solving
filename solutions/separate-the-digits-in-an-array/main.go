package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(separateDigits([]int{13, 25, 83, 77}))
	fmt.Println(separateDigits([]int{7, 1, 3, 9}))
}

func separateDigits(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num < 10 {
			results = append(results, num)
			continue
		}

		vStr := strconv.Itoa(num)
		for j := 0; j < len(vStr); j++ {
			vNum := vStr[j]
			results = append(results, int(vNum-'0'))
		}
	}

	return results
}
