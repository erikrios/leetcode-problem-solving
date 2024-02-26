package main

import "fmt"

func main() {
	fmt.Println(modifiedMatrix([][]int{{1, 2, -1}, {4, -1, 6}, {7, 8, 9}}))
	fmt.Println(modifiedMatrix([][]int{{3, -1}, {5, 2}}))
}

func modifiedMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	maxCols := make([]int, n)
	negOnesIdx := make([]struct{ rowIdx, colIdx int }, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := matrix[i][j]
			if maxCols[j] < num {
				maxCols[j] = num
			}
			if num == -1 {
				negOnesIdx = append(negOnesIdx, struct {
					rowIdx int
					colIdx int
				}{rowIdx: i, colIdx: j})
			}
		}
	}

	for _, val := range negOnesIdx {
		matrix[val.rowIdx][val.colIdx] = maxCols[val.colIdx]
	}

	return matrix
}
