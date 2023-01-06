package main

import "fmt"

func main() {
	fmt.Println(smallerNumberThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumberThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumberThanCurrent([]int{7, 7, 7, 7}))
}

func smallerNumberThanCurrent(nums []int) []int {
	results := make([]int, len(nums))

	for i, num := range nums {
		var counter int

		for j, num2 := range nums {
			if i == j {
				continue
			}

			if num > num2 {
				counter++
			}
		}

		results[i] = counter
	}

	return results
}
