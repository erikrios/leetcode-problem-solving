package main

import "fmt"

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

func longestSubsequence(arr []int, difference int) int {
	mapper := make(map[int]int)
	result := 1

	for _, num := range arr {
		var count int
		if val, ok := mapper[num-difference]; ok {
			count = val
		}

		mapper[num] = count + 1

		if count+1 > result {
			result = count + 1
		}
	}

	return result
}
