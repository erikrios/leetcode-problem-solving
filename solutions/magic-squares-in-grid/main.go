package main

import "fmt"

func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
	fmt.Println(numMagicSquaresInside([][]int{{8}}))
	fmt.Println(numMagicSquaresInside([][]int{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}}))
	fmt.Println(numMagicSquaresInside([][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}}))
}

func numMagicSquaresInside(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	if m < 3 || n < 3 {
		return 0
	}

	var counter int
	for i := 0; i+2 < m; i++ {
		for j := 0; j+2 < n; j++ {
			topLeft := grid[i][j]
			top := grid[i][j+1]
			topRight := grid[i][j+2]
			left := grid[i+1][j]
			center := grid[i+1][j+1]
			right := grid[i+1][j+2]
			bottomLeft := grid[i+2][j]
			bottom := grid[i+2][j+1]
			bottomRight := grid[i+2][j+2]

			square := [3][3]int{
				{topLeft, top, topRight},
				{left, center, right},
				{bottomLeft, bottom, bottomRight},
			}

			if isMagicSquare(square) {
				counter++
			}
		}
	}

	return counter
}

func isMagicSquare(grid [3][3]int) bool {
	m := len(grid)
	uniqueMapper := make(map[int]struct{})

	for _, nums := range grid {
		for _, num := range nums {
			if num < 1 || num > 9 {
				return false
			}
			uniqueMapper[num] = struct{}{}
		}
	}

	if len(uniqueMapper) != m*m {
		return false
	}

	topLeft := grid[0][0]
	top := grid[0][1]
	topRight := grid[0][2]
	left := grid[1][0]
	center := grid[1][1]
	right := grid[1][2]
	bottomLeft := grid[2][0]
	bottom := grid[2][1]
	bottomRight := grid[2][2]

	var sumHorizontals [3]int
	var sumVerticals [3]int
	var sumDiagonals [2]int

	sumHorizontals[0] = topLeft + top + topRight
	sumHorizontals[1] = left + center + right
	sumHorizontals[2] = bottomLeft + bottom + bottomRight

	sumVerticals[0] = topLeft + left + bottomLeft
	sumVerticals[1] = top + center + bottom
	sumVerticals[2] = topRight + right + bottomRight

	sumDiagonals[0] = topLeft + center + bottomRight
	sumDiagonals[1] = topRight + center + bottomLeft

	expected := sumHorizontals[0]

	for _, sumHorizontal := range sumHorizontals {
		if expected != sumHorizontal {
			return false
		}
	}

	for _, sumVertical := range sumVerticals {
		if expected != sumVertical {
			return false
		}
	}

	for _, sumDiagonal := range sumDiagonals {
		if expected != sumDiagonal {
			return false
		}
	}

	return true
}
