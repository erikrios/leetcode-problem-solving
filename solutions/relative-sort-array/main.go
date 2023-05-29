package main

import "fmt"

func main() {
	fmt.Println(relativeSortArray([]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}))
	fmt.Println(relativeSortArray([]int{28, 6, 22, 8, 44, 17}, []int{22, 28, 8, 6}))
}

func relativeSortArray(arr1, arr2 []int) []int {
	n := len(arr1)
	m := len(arr2)
	dist := make(map[int]int, m)
	notExistNums := make([]int, 0, n-m)

	for i := 0; i < m; i++ {
		num := arr2[i]
		dist[num] = 0
	}

	for i := 0; i < n; i++ {
		num := arr1[i]
		if _, ok := dist[num]; ok {
			dist[num]++
		} else {
			notExistNums = append(notExistNums, num)
		}
	}

	results := make([]int, 0, n)
	for i := 0; i < m; i++ {
		num := arr2[i]
		for j := 0; j < dist[num]; j++ {
			results = append(results, num)
		}
	}

	quickSort(notExistNums)

	results = append(results, notExistNums...)

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
