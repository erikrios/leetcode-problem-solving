package main

import "fmt"

func main() {
	fmt.Println(smallestRangeI([]int{1}, 0))
	fmt.Println(smallestRangeI([]int{0, 10}, 2))
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 3))
	fmt.Println(smallestRangeI([]int{7, 8, 8}, 5))
}

func smallestRangeI(nums []int, k int) int {
	min, max := 10000, 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	diff := max - min
	if diff >= -k && diff <= k {
		return 0
	}

	min += k
	max -= k

	if max < min {
		diff := min - max
		if diff <= k {
			max += diff
		} else {
			max += k
		}
	}

	return max - min
}
