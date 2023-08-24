package main

import "fmt"

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

func subsetXORSum(nums []int) int {
	n := len(nums)
	var sum int

	for i := 0; i < (1 << n); i++ {
		var val int

		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				num := nums[j]
				if val == 0 {
					val = num
					continue
				}

				val ^= num
			}
		}

		sum += val
	}

	return sum
}
