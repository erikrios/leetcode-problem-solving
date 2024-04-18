package main

import "fmt"

func main() {
	fmt.Println(resultArray([]int{2, 1, 3}))
	fmt.Println(resultArray([]int{5, 4, 3, 8}))
}

func resultArray(nums []int) []int {
	n := len(nums)
	var addition int
	if n%2 == 1 {
		addition = 1
	}

	arr1 := make([]int, 0, n/2+addition)
	arr2 := make([]int, 0, n/2+addition)

	for i := 0; i < n; i++ {
		num := nums[i]

		if i == 0 {
			arr1 = append(arr1, num)
		} else if i == 1 {
			arr2 = append(arr2, num)
		} else {
			arr1LastElement := arr1[len(arr1)-1]
			arr2LastElement := arr2[len(arr2)-1]

			if arr1LastElement > arr2LastElement {
				arr1 = append(arr1, num)
			} else {
				arr2 = append(arr2, num)
			}
		}
	}

	results := make([]int, 0, n)

	for i := 0; i < len(arr1); i++ {
		results = append(results, arr1[i])
	}

	for i := 0; i < len(arr2); i++ {
		results = append(results, arr2[i])
	}

	return results
}
