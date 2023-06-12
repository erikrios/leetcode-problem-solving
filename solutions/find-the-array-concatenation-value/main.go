package main

import "fmt"

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))
	fmt.Println(findTheArrayConcVal([]int{5, 14, 13, 8, 12}))
}

func findTheArrayConcVal(nums []int) int64 {
	var concat int64

	firstPtr, secondPtr := 0, len(nums)-1
	for firstPtr <= secondPtr {
		if firstPtr == secondPtr {
			concat += int64(nums[firstPtr])
			break
		}

		firstNum := nums[firstPtr]
		secondNum := nums[secondPtr]

		multiplier := 1
		temp := secondNum
		for temp > 0 {
			temp /= 10
			multiplier *= 10
		}

		concat += int64(firstNum*multiplier + secondNum)

		firstPtr++
		secondPtr--
	}

	return concat
}
