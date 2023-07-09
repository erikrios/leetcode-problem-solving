package main

import "fmt"

func main() {
	fmt.Println(
		checkXMatrix([][]int{
			{2, 0, 0, 1},
			{0, 3, 1, 0},
			{0, 5, 2, 0},
			{4, 0, 0, 2},
		}),
	)

	fmt.Println(
		checkXMatrix([][]int{
			{5, 7, 0},
			{0, 3, 1},
			{0, 5, 0},
		}),
	)

	fmt.Println(
		checkXMatrix([][]int{
			{5, 0, 20},
			{0, 5, 0},
			{6, 0, 2},
		}),
	)
}

func checkXMatrix(grid [][]int) bool {
	n := len(grid)
	firstDiagonal := [2]int{0, 0}
	secondDiagonal := [2]int{0, len(grid[0]) - 1}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := grid[i][j]

			if i == firstDiagonal[0] && j == firstDiagonal[1] {
				if num == 0 {
					return false
				}
			} else if i == secondDiagonal[0] && j == secondDiagonal[1] {
				if num == 0 {
					return false
				}
			} else {
				if num != 0 {
					return false
				}
			}
		}

		firstDiagonal[0]++
		firstDiagonal[1]++
		secondDiagonal[0]++
		secondDiagonal[1]--
	}

	return true
}
