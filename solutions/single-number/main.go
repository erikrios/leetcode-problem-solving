package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}

func singleNumber(nums []int) int {
	mapper := make(map[int]bool)

	for _, num := range nums {
		mapper[num] = !mapper[num]
	}

	for k, v := range mapper {
		if v {
			return k
		}
	}

	return -1
}
