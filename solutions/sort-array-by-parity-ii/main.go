package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7}))
	fmt.Println(sortArrayByParityII([]int{2, 3}))
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	for i, j := 0, 1; i <= n-2 && j <= n-1; {
		numI := nums[i]
		numJ := nums[j]
		if numI%2 != 0 && numJ%2 == 0 {
			nums[i], nums[j] = numJ, numI
			numI, numJ = numJ, numI
		}

		if numI%2 == 0 {
			i += 2
		}

		if numJ%2 != 0 {
			j += 2
		}
	}

	return nums
}
