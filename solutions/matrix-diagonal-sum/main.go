package main

import "fmt"

func main() {
	fmt.Println(diagonalSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))

	fmt.Println(diagonalSum([][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}))
}

func diagonalSum(mat [][]int) int {
	var sum int

	fp, sp := 0, len(mat[0])-1
	for i := 0; i < len(mat); i++ {
		fv, sv := mat[i][fp], mat[i][sp]

		if fp == sp {
			sum += fv
		} else {
			sum += fv + sv
		}

		fp++
		sp--
	}

	return sum
}
