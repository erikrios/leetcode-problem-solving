package main

import "fmt"

func main() {
	fmt.Println(smallestEqual([]int{0, 1, 2}))
	fmt.Println(smallestEqual([]int{4, 3, 2, 1}))
	fmt.Println(smallestEqual([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func smallestEqual(nums []int) int {
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if i%10 == num {
			return i
		}
	}

	return -1
}
