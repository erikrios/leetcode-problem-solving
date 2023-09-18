package main

import "fmt"

func main() {
	fmt.Println(
		findColumnWidth(
			[][]int{
				{1},
				{2},
				{333},
			},
		),
	)

	fmt.Println(
		findColumnWidth(
			[][]int{
				{-15, 1, 3},
				{15, 7, 12},
				{5, 6, -2},
			},
		),
	)
}

func findColumnWidth(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var max int

		for j := 0; j < m; j++ {
			num := grid[j][i]
			length := numLen(num)

			if length > max {
				max = length
			}
		}

		res = append(res, max)
	}

	return res
}

func numLen(num int) int {
	var isNegative bool

	if num <= 0 {
		isNegative = true
		num *= -1
	}

	var counter int
	for num > 0 {
		num /= 10
		counter++
	}

	if isNegative {
		counter++
	}

	return counter
}
