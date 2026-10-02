package main

import "fmt"

func main() {
	fmt.Println(findDegrees([][]int{{0, 1, 1}, {1, 0, 1}, {1, 1, 0}}))
	fmt.Println(findDegrees([][]int{{0, 1, 0}, {1, 0, 0}, {0, 0, 0}}))
}

func findDegrees(matrix [][]int) []int {
	results := make([]int, 0, len(matrix))

	for i := 0; i < len(matrix); i++ {
		var res int
		for j := 0; j < len(matrix[i]); j++ {
			res += matrix[i][j]
		}
		results = append(results, res)
	}

	return results
}
