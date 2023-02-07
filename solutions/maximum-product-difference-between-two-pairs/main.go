package main

import "fmt"

func main() {
	fmt.Println(maxProductDifference([]int{5, 6, 2, 7, 4}))
	fmt.Println(maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
}

func maxProductDifference(nums []int) int {
	QuickSort(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[1]*nums[0]
}

func QuickSort(nums []int) {
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
