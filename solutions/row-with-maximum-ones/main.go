package main

import "fmt"

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{{0, 1}, {1, 0}}))
	fmt.Println(rowAndMaximumOnes([][]int{{0, 0, 0}, {0, 1, 1}}))
	fmt.Println(rowAndMaximumOnes([][]int{{0, 0}, {1, 1}, {0, 0}}))
}

func rowAndMaximumOnes(mat [][]int) []int {
	var index int
	var sum int

	for i := 0; i < len(mat); i++ {
		row := mat[i]

		var total int
		for j := 0; j < len(row); j++ {
			total += row[j]
		}

		if total > sum {
			sum = total
			index = i
		}
	}

	return []int{index, sum}
}
