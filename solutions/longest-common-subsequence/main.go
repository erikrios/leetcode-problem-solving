package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
	fmt.Println(longestCommonSubsequence("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func longestCommonSubsequence(text1, text2 string) int {
	m := len(text1)
	n := len(text2)

	matrix := make([][]int, 0, m+1)
	for i := 0; i < cap(matrix); i++ {
		cols := make([]int, n+1)
		matrix = append(matrix, cols)
	}

	for i := len(matrix) - 2; i >= 0; i-- {
		for j := len(matrix[0]) - 2; j >= 0; j-- {
			if text1[i] == text2[j] {
				matrix[i][j] = 1 + matrix[i+1][j+1]
			} else {
				max := matrix[i][j+1]
				if bottom := matrix[i+1][j]; bottom > max {
					max = bottom
				}
				matrix[i][j] = max
			}
		}
	}

	return matrix[0][0]
}
