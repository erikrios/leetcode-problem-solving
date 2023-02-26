package main

import "fmt"

func main() {
	fmt.Println(largestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	}))

	fmt.Println(largestLocal([][]int{
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
}

func largestLocal(grid [][]int) [][]int {
	cap := len(grid) - 2
	results := make([][]int, cap, cap)

	for i := 0; i < cap; i++ {
		rows := make([]int, cap, cap)

		for j := 0; j < cap; j++ {
			max := max(grid, i, j, i+2, j+2)
			rows[j] = max
		}

		results[i] = rows
	}

	return results
}

func max(grid [][]int, startCol, startRow, endCol, endRow int) int {
	var max int

	for i := startCol; i <= endCol; i++ {
		for j := startRow; j <= endRow; j++ {
			val := grid[i][j]
			if max < val {
				max = val
			}
		}
	}

	return max
}
