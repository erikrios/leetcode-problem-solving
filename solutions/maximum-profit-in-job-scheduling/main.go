package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(jobScheduling([]int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}))
	fmt.Println(jobScheduling([]int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}))
	fmt.Println(jobScheduling([]int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}))
	fmt.Println(jobScheduling([]int{6, 15, 7, 11, 1, 3, 16, 2}, []int{19, 18, 19, 16, 10, 8, 19, 8}, []int{2, 9, 1, 19, 5, 7, 3, 19}))
}

func jobScheduling(startTime, endTime, profit []int) int {
	jobs := make([][]int, 0, len(startTime))

	for i := 0; i < len(startTime); i++ {
		job := []int{startTime[i], endTime[i], profit[i]}
		jobs = append(jobs, job)
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	sort.Ints(startTime)

	memo := make([]int, len(startTime))
	for i := range memo {
		memo[i] = -1
	}

	return maxProfit(0, jobs, startTime, memo)
}

func maxProfit(currentIndex int, jobs [][]int, startTime []int, memo []int) int {
	if currentIndex == len(jobs) {
		return 0
	}

	if memo[currentIndex] != -1 {
		return memo[currentIndex]
	}

	nextIndex := sort.Search(len(startTime), func(i int) bool {
		return startTime[i] >= jobs[currentIndex][1]
	})

	// take
	take := jobs[currentIndex][2] + maxProfit(nextIndex, jobs, startTime, memo)

	// not take
	notTake := maxProfit(currentIndex+1, jobs, startTime, memo)

	max := take
	if notTake > max {
		max = notTake
	}

	memo[currentIndex] = max
	return max
}
