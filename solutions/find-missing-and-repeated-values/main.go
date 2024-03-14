package main

import "fmt"

func main() {
	fmt.Println(findMissingAndRepeatedValues([][]int{{1, 3}, {2, 2}}))
	fmt.Println(findMissingAndRepeatedValues([][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}))
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)

	nums := make([]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := grid[i][j]
			nums[val-1]++
		}
	}

	var a, b int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num == 0 {
			b = i + 1
		}

		if num == 2 {
			a = i + 1
		}
	}

	return []int{a, b}
}
