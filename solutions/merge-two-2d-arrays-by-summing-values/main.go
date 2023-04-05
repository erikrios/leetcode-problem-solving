package main

import "fmt"

func main() {
	fmt.Println(mergeArrays([][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 4}, {3, 2}, {4, 1}}))
	fmt.Println(mergeArrays([][]int{{2, 4}, {3, 6}, {5, 5}}, [][]int{{1, 3}, {4, 3}}))
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	var firstPtr, secondPtr int
	firstLength, secondLength := len(nums1), len(nums2)

	results := make([][]int, 0, firstLength+secondLength)
	for firstPtr < firstLength || secondPtr < secondLength {
		res := make([]int, 2, 2)

		if firstPtr < firstLength && secondPtr < secondLength {
			num1 := nums1[firstPtr]
			num2 := nums2[secondPtr]

			if num1[0] == num2[0] {
				res[0], res[1] = num1[0], num1[1]+num2[1]
				firstPtr++
				secondPtr++
			} else if num1[0] < num2[0] {
				res[0], res[1] = num1[0], num1[1]
				firstPtr++
			} else {
				res[0], res[1] = num2[0], num2[1]
				secondPtr++
			}
		} else if firstPtr < firstLength {
			num1 := nums1[firstPtr]
			res[0], res[1] = num1[0], num1[1]
			firstPtr++
		} else {
			num2 := nums2[secondPtr]
			res[0], res[1] = num2[0], num2[1]
			secondPtr++
		}

		results = append(results, res)
	}

	return results
}
