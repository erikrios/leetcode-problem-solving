package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

func minCostClimbingStairs(cost []int) int {
	twoStepBefore := cost[0]
	oneStepBefore := cost[1]
	var curr int

	for i := 2; i < len(cost); i++ {
		val := cost[i]

		var min int
		if twoStepBefore < oneStepBefore {
			min = twoStepBefore
		} else {
			min = oneStepBefore
		}

		curr = min + val

		twoStepBefore = oneStepBefore
		oneStepBefore = curr
	}

	var min int
	if twoStepBefore < oneStepBefore {
		min = twoStepBefore
	} else {
		min = oneStepBefore
	}

	return min
}
