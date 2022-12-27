package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{1, 1, 1, 1, 1}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}

func runningSum(nums []int) []int {
	results := make([]int, len(nums))

	var sum int
	for i, num := range nums {
		sum += num
		results[i] = sum
	}

	return results
}
