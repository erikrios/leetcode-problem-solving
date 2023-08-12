package main

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{
		{1, 2},
		{3, 4},
	}))
	fmt.Println(surfaceArea([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}))
	fmt.Println(surfaceArea([][]int{
		{2, 2, 2},
		{2, 1, 2},
		{2, 2, 2},
	}))
}

func surfaceArea(grid [][]int) int {
	n := len(grid)
	var sum int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]

			var frontV int
			if i-1 >= 0 {
				frontV = grid[i-1][j]
			}

			var backV int
			if i+1 < n {
				backV = grid[i+1][j]
			}

			var leftV int
			if j-1 >= 0 {
				leftV = grid[i][j-1]
			}

			var rightV int
			if j+1 < n {
				rightV = grid[i][j+1]
			}

			topBottomSum := sumTopAndBottom(v)
			frontSum := sumSide(v, frontV)
			backSum := sumSide(v, backV)
			leftSum := sumSide(v, leftV)
			rightSum := sumSide(v, rightV)

			sum += (topBottomSum + frontSum + backSum + leftSum + rightSum)
		}
	}

	return sum
}

func sumTopAndBottom(v int) int {
	if v > 0 {
		return 2
	}

	return 0
}

func sumSide(v, sideV int) int {
	if v <= sideV {
		return 0
	}

	return v - sideV
}
