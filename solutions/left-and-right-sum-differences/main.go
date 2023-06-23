package main

import "fmt"

func main() {
	fmt.Println(leftRigthDifference([]int{10, 4, 8, 3}))
	fmt.Println(leftRigthDifference([]int{1}))
}

func leftRigthDifference(nums []int) []int {
	length := len(nums)
	leftSum, rightSum, answers := make([]int, length, length), make([]int, length, length), make([]int, length, length)

	for i := 0; i < length; i++ {
		if i == 0 {
			leftSum[i] = 0
			rightSum[length-1-i] = 0
			continue
		}

		leftSum[i] = leftSum[i-1] + nums[i-1]
		rightSum[length-1-i] = rightSum[length-1-i+1] + nums[length-1-i+1]
	}

	for i := 0; i < length; i++ {
		v := leftSum[i] - rightSum[i]
		if v < 0 {
			v *= -1
		}
		answers[i] = v
	}

	return answers
}
