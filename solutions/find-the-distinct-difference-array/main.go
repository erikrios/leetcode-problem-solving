package main

import "fmt"

func main() {
	fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(distinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func distinctDifferenceArray(nums []int) []int {
	n := len(nums)
	prefix, suffix := make(map[int]uint8), make(map[int]uint8)

	for i := 0; i < n; i++ {
		num := nums[i]
		suffix[num]++
	}

	results := make([]int, 0, n)
	for i := 0; i < n; i++ {
		num := nums[i]
		prefix[num]++
		suffix[num]--

		if suffix[num] == 0 {
			delete(suffix, num)
		}

		result := len(prefix) - len(suffix)
		results = append(results, result)
	}

	return results
}

func distinctDifferenceArrays(nums []int) []int {
	n := len(nums)
	results := make([]int, 0, n)

	for i := 0; i < n; i++ {
		prefix := countDistinct(nums[:i+1])
		suffix := countDistinct(nums[i+1:])
		result := prefix - suffix
		results = append(results, result)
	}

	return results
}

func countDistinct(nums []int) int {
	n := len(nums)
	numsSet := make(map[int]struct{})
	var counter int

	for i := 0; i < n; i++ {
		num := nums[i]
		if _, ok := numsSet[num]; !ok {
			counter++
			numsSet[num] = struct{}{}
		}
	}

	return counter
}
