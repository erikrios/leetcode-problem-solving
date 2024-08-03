package main

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(1, 0, 0, 0))
	fmt.Println(knightProbability(5, 1, 2, 2))
	fmt.Println(knightProbability(5, 2, 2, 2))
	fmt.Println(knightProbability(3, 2, 0, 0))
}

var directionFormulas = [8][2]int{
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
	{2, -1},
	{2, 1},
}

func knightProbability(n int, k int, row int, column int) float64 {
	dp := make(map[[3]int]float64)
	return backtracking(n, k, row, column, dp)
}

func backtracking(n, k, row, column int, dp map[[3]int]float64) float64 {
	if k == 0 {
		if isInside(row, column, n) {
			return 1
		} else {
			return 0
		}
	}

	if !isInside(row, column, n) {
		return 0
	}

	key := [3]int{k, row, column}
	if val, exists := dp[key]; exists {
		return val
	}

	var probability float64
	for _, dirForm := range directionFormulas {
		rowTarget := row + dirForm[0]
		colTarget := column + dirForm[1]
		probability += backtracking(n, k-1, rowTarget, colTarget, dp) / 8.0
	}

	dp[key] = probability

	return probability
}

func isInside(x, y, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}
