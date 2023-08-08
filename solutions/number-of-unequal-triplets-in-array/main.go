package main

import "fmt"

func main() {
	fmt.Println(unequalTriplets([]int{4, 4, 2, 4, 3}))
	fmt.Println(unequalTriplets([]int{1, 1, 1, 1, 1}))
}

func unequalTriplets(nums []int) int {
	n := len(nums)

	var counter int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				numI := nums[i]
				numJ := nums[j]
				numK := nums[k]

				if numI != numJ && numI != numK && numJ != numK {
					counter++
				}
			}
		}
	}

	return counter
}
