package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(matrixScore([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
	fmt.Println(matrixScore([][]int{{0}}))
	fmt.Println(matrixScore([][]int{{0, 1, 0}, {0, 0, 1}}))
}

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			flipMatrixRow(i, grid)
		}
	}

	for j := 0; j < n; j++ {
		var oneCounter int
		for i := 0; i < m; i++ {
			oneCounter += grid[i][j]
		}
		if oneCounter < m-oneCounter {
			flipMatrixCol(j, grid)
		}
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res += int(math.Pow(2.0, float64(n-1-j)))
			}
		}
	}

	return res
}

func flipMatrixRow(i int, grid [][]int) {
	for j := 0; j < len(grid[i]); j++ {
		cur := grid[i][j]
		if cur == 1 {
			grid[i][j] = 0
		} else {
			grid[i][j] = 1
		}
	}
}

func flipMatrixCol(j int, grid [][]int) {
	m := len(grid)

	for i := 0; i < m; i++ {
		cur := grid[i][j]
		if cur == 1 {
			grid[i][j] = 0
		} else {
			grid[i][j] = 1
		}
	}
}
