package main

import "fmt"

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
	fmt.Println(processQueries([]int{4, 1, 2, 2}, 4))
	fmt.Println(processQueries([]int{7, 5, 5, 8, 3}, 8))
}

func processQueries(queries []int, m int) []int {
	n := len(queries)
	perms := generatePermutations(m)
	results := make([]int, 0, n)

	for i := 0; i < n; i++ {
		query := queries[i]
		pos := idxPos(perms, query)
		moveToBeginning(perms, pos)
		results = append(results, pos)
	}

	return results
}

func generatePermutations(n int) []int {
	res := make([]int, 0, n)

	for i := 0; i < n; i++ {
		res = append(res, i+1)
	}

	return res
}

func idxPos(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func moveToBeginning(nums []int, index int) {
	temp := nums[index]

	for i := index; i > 0; i-- {
		nums[i] = nums[i-1]
	}

	nums[0] = temp
}
