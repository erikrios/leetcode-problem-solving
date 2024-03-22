package main

import "fmt"

func main() {
	fmt.Println(maximumStrongPairXor([]int{1, 2, 3, 4, 5}))
	fmt.Println(maximumStrongPairXor([]int{10, 100}))
	fmt.Println(maximumStrongPairXor([]int{5, 6, 25, 30}))
}

func maximumStrongPairXor(nums []int) int {
	var maxXor int

	for i := 0; i < len(nums); i++ {
		numOne := nums[i]
		for j := i; j < len(nums); j++ {
			numTwo := nums[j]

			diff := numOne - numTwo
			if diff < 0 {
				diff *= -1
			}

			min := numOne
			if numTwo < numOne {
				min = numTwo
			}

			if isStrongPair := diff <= min; isStrongPair {
				xorResult := numOne ^ numTwo
				if xorResult > maxXor {
					maxXor = xorResult
				}
			}
		}
	}

	return maxXor
}
