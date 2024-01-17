package main

import "fmt"

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{2, 4, 6, 8, 10}))
	fmt.Println(numberOfArithmeticSlices([]int{7, 7, 7, 7, 7}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(numberOfArithmeticSlices([]int{1, 2, 1, 2, 4, 1, 5, 10}))
	fmt.Println(numberOfArithmeticSlices([]int{0, 2, 2, 3, 4}))
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	var ans int

	dp := make([]map[int]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			var cnt int
			if v, ok := dp[j][diff]; ok {
				cnt = v
			}

			var temp int
			if v, ok := dp[i][diff]; ok {
				temp = v
			}

			dp[i][diff] = temp + cnt + 1
			ans += cnt
		}
	}

	return ans
}
