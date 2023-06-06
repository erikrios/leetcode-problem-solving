package main

import "fmt"

func main() {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
	fmt.Println(minimumOperations([]int{1, 5, 2, 3, 5}))
	fmt.Println(minimumOperations([]int{1, 5, 3, 3, 5}))
	fmt.Println(minimumOperations([]int{0}))
}

func minimumOperations(nums []int) int {
	quickSort(nums)

	var counter int

	var starter int
	for nums[len(nums)-1] > 0 {
		var min int

		for i := starter; i < len(nums); i++ {
			num := nums[i]
			if num == 0 {
				continue
			}

			if min == 0 {
				min = num
			}

			nums[i] = num - min
			if nums[i] == 0 {
				starter = i + 1
			}
		}

		counter++
	}

	return counter
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
