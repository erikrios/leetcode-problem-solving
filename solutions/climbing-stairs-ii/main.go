package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(climbStairs(4, []int{1, 2, 3, 4}))
	fmt.Println(climbStairs(4, []int{5, 1, 6, 2}))
	fmt.Println(climbStairs(3, []int{9, 8, 3}))
	fmt.Println(climbStairs(7, []int{11, 10, 2, 10, 11, 9, 2}))
}

func climbStairs(n int, costs []int) int {
	return dfs(n, costs, 0, make(map[int]int))
}

func dfs(n int, costs []int, i int, memo map[int]int) int {
	if i == n {
		return 0
	}

	if val, ok := memo[i]; ok {
		return val
	}

	var minCost = math.MaxInt
	for j := i + 1; j <= i+3 && j <= n; j++ {
		cost := costs[j-1] + (j-i)*(j-i)
		total := cost + dfs(n, costs, j, memo)
		if total < minCost {
			minCost = total
		}
	}

	memo[i] = minCost

	return minCost
}
