package main

import "fmt"

func main() {
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 2))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 3))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 5))
}

func targetIndices(nums []int, target int) []int {
	quickSort(nums)

	results := make([]int, 0)
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		num := nums[m]

		if num < target {
			l = m + 1
		} else if num > target {
			r = m - 1
		} else {
			results = append(results, m)
			lIndex, rIndex := m-1, m+1

			for {
				var atLeastOneMet bool

				if lIndex >= 0 && num == nums[lIndex] {
					results = append(results, 0)
					for i := len(results) - 1; i > 0; i-- {
						results[i] = results[i-1]
					}
					results[0] = lIndex
					lIndex--
					atLeastOneMet = true

				}

				if rIndex < len(nums) && num == nums[rIndex] {
					results = append(results, rIndex)
					rIndex++
					atLeastOneMet = true
				}

				if !atLeastOneMet {
					break
				}
			}

			break
		}
	}

	return results
}

func quickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	sort(nums, 0, len(nums)-1)
}

func sort(nums []int, startIndex, endIndex int) {
	if endIndex-startIndex <= 0 {
		return
	}

	p := partition(nums, startIndex, endIndex)

	sort(nums, startIndex, p-1)
	sort(nums, p+1, endIndex)
}

func partition(nums []int, startIndex, endIndex int) int {
	j := startIndex - 1

	for i := startIndex; i < endIndex; i++ {
		if nums[i] < nums[endIndex] {
			j++
			if i == j {
				continue
			}
			swap(nums, j, i)
		}
	}

	swap(nums, j+1, endIndex)

	return j + 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
