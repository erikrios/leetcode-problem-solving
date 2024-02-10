package main

import (
	"fmt"
)

func main() {
	fmt.Println(onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}))
	fmt.Println(onesMinusZeros([][]int{{1, 1, 1}, {1, 1, 1}}))
}

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	oneRows := make([]int, m)
	oneCols := make([]int, n)

	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				oneRows[i]++
				oneCols[j]++
			}
		}
	}

	for i := range grid {
		for j := range grid[0] {
			grid[i][j] = oneRows[i] + oneCols[j] - (m - oneRows[i]) - (n - oneCols[j])
		}
	}

	return grid
}
