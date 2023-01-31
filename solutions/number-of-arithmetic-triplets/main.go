package main

import "fmt"

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
}

func arithmeticTriplets(nums []int, diff int) int {
	var total int

	for i := 1; i <= len(nums)-1-1; i++ {
		leftIndex, rightIndex := i-1, i+1

		for leftIndex >= 0 && rightIndex <= len(nums)-1 {
			leftDiff := nums[i] - nums[leftIndex]
			rightDiff := nums[rightIndex] - nums[i]

			if leftDiff == diff && rightDiff == diff {
				total++
				break
			}

			if leftDiff < diff {
				leftIndex--
			} else if leftDiff > diff {
				break
			}

			if rightDiff < diff {
				rightIndex++
			} else if rightDiff > diff {
				break
			}
		}
	}

	return total
}
