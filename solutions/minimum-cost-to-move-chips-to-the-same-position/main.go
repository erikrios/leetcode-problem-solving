package main

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	fmt.Println(minCostToMoveChips([]int{1, 1_000_000_000}))
	fmt.Println(minCostToMoveChips([]int{6, 4, 7, 8, 2, 10, 2, 7, 9, 7}))
}

func minCostToMoveChips(position []int) int {
	var posCount, negCount int

	for i := 0; i < len(position); i++ {
		pos := position[i]
		if pos%2 == 0 {
			posCount++
		} else {
			negCount++
		}
	}

	if posCount > negCount {
		return negCount
	}

	return posCount
}
