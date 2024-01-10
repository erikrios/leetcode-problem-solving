package main

import "fmt"

func main() {
	fmt.Println(numberGame([]int{5, 4, 2, 3}))
	fmt.Println(numberGame([]int{2, 5}))
}

func numberGame(nums []int) []int {
	bubbleSort(nums)

	for i := 0; i < len(nums); i += 2 {
		a, b := nums[i], nums[i+1]
		if b > a {
			nums[i], nums[i+1] = b, a
		}
	}

	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			currentNum := nums[i]
			nextNum := nums[j]
			if nextNum < currentNum {
				nums[i], nums[j] = nextNum, currentNum
			}
		}
	}
}
