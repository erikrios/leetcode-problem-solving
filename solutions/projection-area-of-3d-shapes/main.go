package main

import "fmt"

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}))
	fmt.Println(projectionArea([][]int{{2}}))
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}}))
}

func projectionArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var xyAxis, xzAxis, yzAxis int
	maximums := make([]int, n, n)

	for i := 0; i < m; i++ {
		nums := grid[i]

		var max int
		for j := 0; j < n; j++ {
			num := nums[j]

			if num > 0 {
				xyAxis++
			}

			if num > max {
				max = num
			}

			if num > maximums[j] {
				maximums[j] = num
			}
		}

		xzAxis += max
	}

	for i := 0; i < len(maximums); i++ {
		yzAxis += maximums[i]
	}

	return xyAxis + xzAxis + yzAxis
}
