package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}

func decompressRLElist(nums []int) []int {
	results := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			results = append(results, val)
		}
	}

	return results
}
