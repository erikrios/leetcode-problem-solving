package main

import "fmt"

func main() {
	fmt.Println(hasTrailingZeros([]int{1, 2, 3, 4, 5}))
	fmt.Println(hasTrailingZeros([]int{2, 4, 8, 16}))
	fmt.Println(hasTrailingZeros([]int{1, 3, 5, 7, 9}))
}

func hasTrailingZeros(nums []int) bool {
	var hasTrailingZeros bool

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num&1 != 1 {
			if hasTrailingZeros {
				return true
			}

			hasTrailingZeros = true
		}
	}

	return false
}
