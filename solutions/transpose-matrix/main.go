package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(transpose([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		nums := make([]int, 0, m)
		for j := 0; j < m; j++ {
			num := matrix[j][i]
			nums = append(nums, num)
		}

		res = append(res, nums)
	}

	return res
}
