package main

import "fmt"

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
	fmt.Println(arraySign([]int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}))
}

func arraySign(nums []int) int {
	res := 1

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num == 0 {
			return 0
		}

		if num < 0 {
			res = -res
		}
	}

	return res
}
