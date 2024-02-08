package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(10_000))
}

func numSquares(n int) int {
	return numSquaresMemo(n, make(map[int]int))
}

func numSquaresMemo(n int, memo map[int]int) int {
	if n == 0 {
		return 0
	}

	if v, ok := memo[n]; ok {
		return v
	}

	min := math.MaxInt
	for i := 1; i*i <= n; i++ {
		if val := numSquaresMemo(n-i*i, memo); val < min {
			min = val
		}
	}

	min++
	memo[n] = min
	return min
}
