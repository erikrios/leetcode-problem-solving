package main

import "fmt"

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
