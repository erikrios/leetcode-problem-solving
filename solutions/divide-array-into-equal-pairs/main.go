package main

import "fmt"

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
	fmt.Println(divideArray([]int{1, 2, 3, 4}))
}

func divideArray(nums []int) bool {
	mapper := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if _, ok := mapper[num]; ok {
			delete(mapper, num)
		} else {
			mapper[num] = true
		}
	}

	return len(mapper) == 0
}
