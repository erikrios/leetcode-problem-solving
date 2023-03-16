package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{1,2,3,3}))
	fmt.Println(repeatedNTimes([]int{2,1,2,5,3,2}))
	fmt.Println(repeatedNTimes([]int{5,1,5,2,5,3,5,4}))
}

func repeatedNTimes(nums []int) int {
	n := len(nums)/2
	numCounts := make(map[int]int)

	var res int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		numCounts[num]++
		if numCounts[num] == n {
			res = num
		}
	}

	return res
}
