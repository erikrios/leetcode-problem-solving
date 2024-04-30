package main

import "fmt"

func main() {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
	fmt.Println(decrypt([]int{1, 2, 3, 4}, 0))
	fmt.Println(decrypt([]int{2, 4, 9, 3}, -2))
}

func decrypt(code []int, k int) []int {
	n := len(code)
	results := make([]int, n)

	if k == 0 {
		return results
	}

	isNeg := k < 0

	start, end := 1, k

	if isNeg {
		start, end = n+k, n-1
	}

	sum := sum(start, end, code)

	for i := 0; i < n; i++ {
		results[i] = sum

		prev := start

		start++
		end++

		if start >= n {
			start = 0
		}

		if end >= n {
			end = 0
		}

		sum -= code[prev]
		sum += code[end]
	}

	return results
}

func sum(start, end int, nums []int) int {
	var sum int
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	return sum
}
