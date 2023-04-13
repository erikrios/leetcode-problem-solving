package main

import "fmt"

func main() {
	fmt.Println(countNegatives([][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}))
	fmt.Println(countNegatives([][]int{{3, 2}, {1, 0}}))
	fmt.Println(countNegatives([][]int{{5, 1, 0}, {-5, -5, -5}}))
}

func countNegative(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var counter int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			if v < 0 {
				counter += (n - j)
				break
			}
		}
	}

	return counter
}

func countNegatives(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	var counter int
	for i := 0; i < m; i++ {
		values := grid[i]
		l, r := 0, n-1
		for l <= r {
			mid := (l + r) / 2
			v := values[mid]
			if v > -1 {
				l = mid + 1
			} else {
				if mid > 0 && values[mid-1] <= -1 {
					l--
					r--
					continue
				}
				counter += (n - mid)
				break
			}
		}
	}

	return counter
}
