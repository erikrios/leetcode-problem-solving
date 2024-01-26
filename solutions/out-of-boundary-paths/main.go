package main

import "fmt"

func main() {
	fmt.Println(findPaths(2, 2, 2, 0, 0))
	fmt.Println(findPaths(1, 3, 3, 0, 1))
	fmt.Println(findPaths(8, 7, 16, 1, 5))
	fmt.Println(findPaths(36, 5, 50, 15, 3))
}

func findPaths(m, n, maxMove, startRow, startColumn int) int {
	memo := make([][][]int, m)
	for i := range memo {
		memo[i] = make([][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([]int, maxMove+1)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	return findPath(m, n, maxMove, startRow, startColumn, memo)
}

func findPath(m, n, maxMove, startRow, startColumn int, memo [][][]int) int {
	if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n {
		return 1
	}

	if maxMove <= 0 {
		return 0
	}

	if v := memo[startRow][startColumn][maxMove]; v != -1 {
		return v
	}

	maxMove--
	startRow--
	goTopRes := findPath(m, n, maxMove, startRow, startColumn, memo)
	maxMove++
	startRow++

	maxMove--
	startRow++
	goDownRes := findPath(m, n, maxMove, startRow, startColumn, memo)
	maxMove++
	startRow--

	maxMove--
	startColumn--
	goLeftRes := findPath(m, n, maxMove, startRow, startColumn, memo)
	maxMove++
	startColumn++

	maxMove--
	startColumn++
	goRightRes := findPath(m, n, maxMove, startRow, startColumn, memo)
	maxMove++
	startColumn--

	res := (goTopRes + goDownRes + goLeftRes + goRightRes) % (1_000_000_000 + 7)
	memo[startRow][startColumn][maxMove] = res
	return res
}
