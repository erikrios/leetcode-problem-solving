package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	mapper := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		remaining := target - num
		if idx, ok := mapper[remaining]; ok {
			return []int{idx, i}
		}

		mapper[num] = i
	}

	return []int{}
}
