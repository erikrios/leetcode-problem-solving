package main

import "fmt"

func main() {
	fmt.Println(maxIncreaseKeepingSkyline(
		[][]int{
			{3, 0, 8, 4},
			{2, 4, 5, 7},
			{9, 2, 6, 3},
			{0, 3, 1, 0},
		},
	))
	fmt.Println(maxIncreaseKeepingSkyline(
		[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	))
}

type existence struct {
	ok  bool
	val int
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])

	maxRows := make([]existence, n, n)
	maxCols := make([]existence, m, m)

	var sum int
	for i := 0; i < n; i++ {
		maxRow := maxRows[i]
		if !maxRow.ok {
			maxRow.val = maximumRow(grid[i])
			maxRow.ok = true
			maxRows[i] = maxRow
		}

		for j := 0; j < m; j++ {
			maxCol := maxCols[j]
			if !maxCol.ok {
				maxCol.val = maximumCol(grid, j)
				maxCol.ok = true
				maxCols[j] = maxCol
			}

			var min int

			if maxRow.val < maxCol.val {
				min = maxRow.val
			} else {
				min = maxCol.val
			}

			height := grid[i][j]
			if height < min {
				diff := min - height
				sum += diff
				grid[i][j] = min
			}
		}
	}

	return sum
}

func maximumRow(rows []int) int {
	var max int

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if max < row {
			max = row
		}
	}

	return max
}

func maximumCol(grid [][]int, j int) int {
	var max int

	for i := 0; i < len(grid); i++ {
		rows := grid[i]
		col := rows[j]
		if max < col {
			max = col
		}
	}

	return max
}
