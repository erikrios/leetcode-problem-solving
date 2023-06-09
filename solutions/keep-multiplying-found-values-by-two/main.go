package main

import "fmt"

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
	fmt.Println(findFinalValue([]int{2, 7, 9}, 4))
	fmt.Println(findFinalValue([]int{8, 19, 4, 2, 15, 3}, 2))
}

func findFinalValue(nums []int, original int) int {
	quickSort(nums)

	for startIndex := 0; ; {
		startIndex = binarySearch(nums[startIndex:], original)

		if startIndex == -1 {
			break
		}

		original *= 2
	}

	return original
}

func binarySearch(nums []int, target int) int {
	index := -1

	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			index = mid
			break
		}
	}

	return index
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
