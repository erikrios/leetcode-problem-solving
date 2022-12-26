package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4, 5, 6}, 3))
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	fmt.Println(shuffle([]int{1, 1, 2, 2}, 2))
}

func shuffle(nums []int, n int) []int {
	results := make([]int, len(nums))

	lp, rp := 0, n

	for i := 0; rp < len(nums); i += 2 {
		results[i] = nums[lp]
		results[i+1] = nums[rp]

		lp++
		rp++
	}

	return results
}
