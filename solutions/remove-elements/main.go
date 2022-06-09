package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			nums = append(nums[0:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
