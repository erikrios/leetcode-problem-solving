package main

import "fmt"

func main() {
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}))
	fmt.Println(numSpecial([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	fmt.Println(numSpecial([][]int{
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 0, 0},
	}))
}

func numSpecial(mat [][]int) int {
	rowsTrackers := make([]int, 0, len(mat))
	colsSum := make([]int, len(mat[0]), len(mat[0]))

	for i := 0; i < len(mat); i++ {
		var rowSum int
		var colIdx int
		for j := 0; j < len(mat[i]); j++ {
			num := mat[i][j]
			rowSum += num
			colsSum[j] += num
			if num == 1 {
				colIdx = j
			}
		}

		if rowSum == 1 {
			rowsTrackers = append(rowsTrackers, colIdx)
		}
	}

	var counter int
	for i := 0; i < len(rowsTrackers); i++ {
		if colIdx := rowsTrackers[i]; colsSum[colIdx] == 1 {
			counter++
		}
	}

	return counter
}
