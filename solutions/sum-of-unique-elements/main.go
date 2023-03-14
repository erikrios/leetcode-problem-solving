package main

import "fmt"

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
	fmt.Println(sumOfUnique([]int{1, 1, 1, 1, 1}))
	fmt.Println(sumOfUnique([]int{1, 2, 3, 4, 5}))
}

func sumOfUnique(nums []int) int {
	mapper := make(map[int]int)
	var sum int

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if v, ok := mapper[num]; !ok {
			sum += num
		} else if v == 1 {
			sum -= num

		}

		mapper[num]++
	}

	return sum
}
