package main

import "fmt"

func main() {
	fmt.Println(sumCounts([]int{1, 2, 1}))
	fmt.Println(sumCounts([]int{1, 1}))
}

func sumCounts(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		for j := 0; j+i+1 <= len(nums); j++ {
			results := nums[j : j+i+1]
			distinctMap := make(map[int]struct{})
			for _, val := range results {
				distinctMap[val] = struct{}{}
			}

			distinct := len(distinctMap)
			sum += (distinct * distinct)
		}
	}

	return sum
}
