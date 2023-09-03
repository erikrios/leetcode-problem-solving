package main

import "fmt"

func main() {
	fmt.Println(sortEvenOdd([]int{4, 1, 2, 3}))
	fmt.Println(sortEvenOdd([]int{2, 1}))
	fmt.Println(sortEvenOdd([]int{4, 1, 5, 3, 2}))
}

func sortEvenOdd(nums []int) []int {
	n := len(nums)

	if n <= 2 {
		return nums
	}

	for i := 0; i < n-2; i += 2 {
		curEven := nums[i]
		curOdd := nums[i+1]

		for j := i + 2; j < n; j += 2 {
			nextEven := nums[j]
			if curEven > nextEven {
				nums[i], nums[j] = nextEven, curEven
				curEven = nextEven
			}

			if j+1 >= n {
				continue
			}

			nextOdd := nums[j+1]
			if curOdd < nextOdd {
				nums[i+1], nums[j+1] = nums[j+1], nums[i+1]
				curOdd = nextOdd
			}
		}
	}

	return nums
}
