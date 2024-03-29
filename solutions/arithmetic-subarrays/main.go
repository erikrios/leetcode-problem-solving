package main

import "sort"

func main() {
}

func checkArithmeticSubarrays(nums, l, r []int) []bool {
	m := len(r)
	results := make([]bool, m)

loop:
	for i := 0; i < m; i++ {
		start := l[i]
		end := r[i]

		if end-start < 1 {
			results[i] = false
			continue
		}

		subNum := nums[start : end+1]
		sub := make([]int, len(subNum))
		copy(sub, subNum)
		sort.Ints(sub)

		diff := sub[1] - sub[0]
		for j := 1; j < len(sub)-1; j++ {
			cur := sub[j]
			next := sub[j+1]
			if next-cur != diff {
				results[i] = false
				continue loop
			}
		}
		results[i] = true
	}

	return results
}
