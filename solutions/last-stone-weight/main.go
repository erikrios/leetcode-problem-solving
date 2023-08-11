package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{1}))
	fmt.Println(lastStoneWeight([]int{2, 2}))
}

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		n := len(stones)
		last := stones[n-1]
		prev := stones[n-1-1]

		if prev < last {
			stones[n-1-1] = last - prev
			stones = stones[:n-1]
		} else {
			stones = stones[:n-1-1]
		}
	}

	if len(stones) == 0 {
		return 0
	}

	return stones[0]
}
