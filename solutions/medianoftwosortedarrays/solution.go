package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sortedArray := make([]int, 0)
	sortedArray = append(sortedArray, nums1...)
	sortedArray = append(sortedArray, nums2...)

	for i := 0; i < len(sortedArray)-1; i++ {
		for j := 0; j < len(sortedArray)-i-1; j++ {
			if sortedArray[j] > sortedArray[j+1] {
				temp := sortedArray[j]
				sortedArray[j] = sortedArray[j+1]
				sortedArray[j+1] = temp
			}
		}
	}

	var result float64

	if len(sortedArray)%2 == 0 {
		middleOne := sortedArray[len(sortedArray)/2-1]
		middleTwo := sortedArray[len(sortedArray)/2]

		result = (float64(middleOne) + float64(middleTwo)) / 2
	} else {
		result = float64(sortedArray[len(sortedArray)/2])
	}

	return result
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
