package main

import "fmt"

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
}

func arrayPairSum(nums []int) int {
	quickSort(nums)
	var maxSum int

	for i := 0; i < len(nums); i += 2 {
		maxSum += nums[i]
	}

	return maxSum
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
