package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for iStart, jStart, iEnd, jEnd := 0, 0, n-1, n-1; iStart < iEnd; iStart, jStart, iEnd, jEnd = iStart+1, jStart+1, iEnd-1, jEnd-1 {
		rotation(matrix, iStart, jStart, iEnd, jEnd)
	}
}

func rotation(matrix [][]int, iStart, jStart, iEnd, jEnd int) {
	times := iEnd - iStart

	for i := 0; i < times; i++ {
		tmp := matrix[iStart][jStart]

		for left := iStart; left < iEnd; left++ {
			matrix[left][jStart] = matrix[left+1][jStart]
		}

		for bottom := jStart; bottom < jEnd; bottom++ {
			matrix[iEnd][bottom] = matrix[iEnd][bottom+1]
		}

		for right := iEnd; right > iStart; right-- {
			matrix[right][jEnd] = matrix[right-1][jEnd]
		}

		for top := jEnd; top > jStart; top-- {
			matrix[iStart][top] = matrix[iStart][top-1]
		}

		matrix[iStart][jStart+1] = tmp
	}
}
