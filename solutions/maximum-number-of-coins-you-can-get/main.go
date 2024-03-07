package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(maxCoins([]int{2, 4, 5}))
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}

func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)

	var counter int
	for i := n / 3; i < n; i += 2 {
		myCoin := piles[i]
		counter += myCoin
	}

	return counter
}
