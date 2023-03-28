package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity([]int{0}))
}

func sortArrayByParity(nums []int) []int {
	results := make([]int, len(nums), len(nums))
	i, j := 0, len(nums)-1

	for l := 0; l < len(nums); l++ {
		num := nums[l]

		if num%2 == 1 {
			results[j] = num
			j--
		} else {
			results[i] = num
			i++
		}
	}

	return results
}
