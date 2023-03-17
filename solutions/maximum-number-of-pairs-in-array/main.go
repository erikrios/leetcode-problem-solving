package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1,3,2,1,3,2,2}))
	fmt.Println(numberOfPairs([]int{1,1}))
	fmt.Println(numberOfPairs([]int{0}))
}

func numberOfPairs(nums []int) []int {
	mapper := make(map[int]uint8)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		mapper[num]++
	}

	results := make([]int, 2, 2)

	for _, v := range mapper {
		results[0] += (int(v) / 2)
		results[1] += (int(v) % 2)
	}

	return results
}
