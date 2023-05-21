package main

import "fmt"

func main() {
	fmt.Println(isToeplitzMatrix([][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}))
	fmt.Println(isToeplitzMatrix([][]int{{1, 2}, {2, 2}}))
}

func isToeplitzMatrix(matrix [][]int) bool {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			compRow := i - 1
			compCol := j - 1

			if compRow < 0 || compCol < 0 {
				continue
			}

			num := matrix[i][j]
			compNum := matrix[compRow][compCol]

			if num != compNum {
				return false
			}
		}
	}

	return true
}
