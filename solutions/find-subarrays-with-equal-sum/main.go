package main

import "fmt"

func main() {
	fmt.Println(findSubarrays([]int{4, 2, 4}))
	fmt.Println(findSubarrays([]int{1, 2, 3, 4, 5}))
	fmt.Println(findSubarrays([]int{0, 0, 0}))
}

func findSubarrays(nums []int) bool {
	mapper := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		num := nums[i]
		next := nums[i+1]
		sum := num + next

		mapper[sum]++

		if mapper[sum] == 2 {
			return true
		}
	}

	return false
}
