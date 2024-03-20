package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(diagonalSort([][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}))
	fmt.Println(diagonalSort([][]int{{11, 25, 66, 1, 69, 7}, {23, 55, 17, 45, 15, 52}, {75, 31, 36, 44, 58, 8}, {22, 27, 33, 25, 68, 4}, {84, 28, 14, 11, 5, 50}}))
}

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	for colIdx := 0; colIdx < n; colIdx++ {
		tempRowIdx := 0
		tempColIdx := colIdx

		values := make([]int, 0)
		for tempRowIdx < m && tempColIdx < n {
			values = append(values, mat[tempRowIdx][tempColIdx])
			tempRowIdx++
			tempColIdx++
		}

		sort.Ints(values)

		tempRowIdx = 0
		tempColIdx = colIdx
		for i := 0; i < len(values); i++ {
			mat[tempRowIdx][tempColIdx] = values[i]
			tempRowIdx++
			tempColIdx++
		}
	}

	for rowIdx := 1; rowIdx < m; rowIdx++ {
		tempRowIdx := rowIdx
		tempColIdx := 0

		values := make([]int, 0)
		for tempRowIdx < m && tempColIdx < n {
			values = append(values, mat[tempRowIdx][tempColIdx])
			tempRowIdx++
			tempColIdx++
		}

		sort.Ints(values)

		tempRowIdx = rowIdx
		tempColIdx = 0
		for i := 0; i < len(values); i++ {
			mat[tempRowIdx][tempColIdx] = values[i]
			tempRowIdx++
			tempColIdx++
		}
	}

	return mat
}
