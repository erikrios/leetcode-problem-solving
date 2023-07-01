package main

import "fmt"

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
	fmt.Println(findTheDistanceValue([]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3))
	fmt.Println(findTheDistanceValue([]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6))
}

func findTheDistanceValue(arr1, arr2 []int, d int) int {
	var counter int

	for i := 0; i < len(arr1); i++ {
		num1 := arr1[i]

		isValid := true
		for j := 0; j < len(arr2); j++ {
			num2 := arr2[j]

			var dist int

			if num1 > num2 {
				dist = num1 - num2
			} else {
				dist = num2 - num1
			}

			if dist <= d {
				isValid = false
				break
			}
		}

		if isValid {
			counter++
		}
	}

	return counter
}
