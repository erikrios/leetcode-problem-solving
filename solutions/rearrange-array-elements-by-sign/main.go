package main

import (
	"fmt"
)

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	fmt.Println(rearrangeArray([]int{-1, 1}))
}

func rearrangeArray(nums []int) []int {
	results := make([]int, len(nums))

	posIdx, negIdx := 0, 1
	for _, num := range nums {
		if num > 0 {
			results[posIdx] = num
			posIdx += 2
		} else {
			results[negIdx] = num
			negIdx += 2
		}
	}

	return results
}
