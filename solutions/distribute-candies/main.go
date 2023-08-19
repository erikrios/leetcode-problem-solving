package main

import "fmt"

func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(distributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(distributeCandies([]int{6, 6, 6, 6}))
}

func distributeCandies(candyType []int) int {
	types := make(map[int]struct{})

	n := len(candyType)
	max := n / 2

	for i := 0; i < n; i++ {
		types[candyType[i]] = struct{}{}
		if len(types) >= max {
			return max
		}
	}

	return len(types)
}
