package main

import "fmt"

func main() {
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
	fmt.Println(answerQueries([]int{2, 3, 4, 5}, []int{1}))
}

func answerQueries(nums []int, queries []int) []int {
	quickSort(nums)
	results := make([]int, len(queries), len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]

		var sum int
		var counter int
		for j := 0; j < len(nums); j++ {
			sum += nums[j]
			if query < sum {
				break
			}

			counter++
			if query == sum {
				break
			}
		}

		results[i] = counter
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
