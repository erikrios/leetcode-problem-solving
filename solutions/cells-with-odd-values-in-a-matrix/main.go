package main

import "fmt"

func main() {
	fmt.Println(oddCells(2, 3, [][]int{{0, 1}, {1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{{1, 1}, {0, 0}}))
}

func oddCells(m int, n int, indices [][]int) int {
	matrix := make([][]int, m, m)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, n, n)
	}

	for i := 0; i < len(indices); i++ {
		rowIndex := indices[i][0]
		colIndex := indices[i][1]

		for j := 0; j < n; j++ {
			matrix[rowIndex][j]++
		}

		for j := 0; j < m; j++ {
			matrix[j][colIndex]++
		}
	}

	var counter int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j]%2 != 0 {
				counter++
			}
		}
	}

	return counter
}
