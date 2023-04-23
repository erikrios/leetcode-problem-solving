package main

import "fmt"

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
	fmt.Println(minSubsequence([]int{4, 4, 7, 6, 7}))
}

func minSubsequence(nums []int) []int {
	quickSort(nums)

	var results []int
	for i := 1; i <= len(nums); i++ {
		inc := nums[:i]
		exc := nums[i:]

		if sum(inc) > sum(exc) {
			results = inc
			break
		}
	}

	return results
}

func sum(nums []int) int {
	var total int

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		total += num
	}

	return total
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
		if nums[i] > nums[endIndex] {
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
