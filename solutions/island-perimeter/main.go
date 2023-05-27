package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(islandPerimeter([][]int{{1}}))
	fmt.Println(islandPerimeter([][]int{{1, 0}}))
}

func islandPerimeter(grid [][]int) int {
	var perim int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if v := grid[i][j]; v == 0 {
				continue
			}

			// left side
			if j-1 < 0 || grid[i][j-1] == 0 {
				perim++
			}

			// top side
			if i-1 < 0 || grid[i-1][j] == 0 {
				perim++
			}

			// right side
			if j+1 > len(grid[i])-1 || grid[i][j+1] == 0 {
				perim++
			}

			// bottom side
			if i+1 > len(grid)-1 || grid[i+1][j] == 0 {
				perim++
			}
		}
	}

	return perim
}
