package main

import "fmt"

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
}

func search(nums []int, target int) int {
	for low, high := 0, len(nums)-1; low <= high; {
		mid := (low + high) / 2

		if num := nums[mid]; num < target {
			low = mid + 1
		} else if num > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
