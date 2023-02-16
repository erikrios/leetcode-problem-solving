package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
	fmt.Println(maxProduct([]int{5, 4, 3, 2}))
}

func maxProducts(nums []int) int {
	var fMaxI int
	sMaxI := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[fMaxI] {
			temp := fMaxI
			fMaxI = i

			if fMaxI == sMaxI {
				sMaxI = i - 1
				continue
			}

			if nums[fMaxI]-nums[sMaxI] > nums[fMaxI]-nums[temp] {
				sMaxI = temp
			}
			continue
		}

		if nums[fMaxI]-nums[sMaxI] > nums[fMaxI]-nums[i] {
			sMaxI = i
		}
	}

	return (nums[fMaxI] - 1) * (nums[sMaxI] - 1)
}

func maxProduct(nums []int) int {
	var fMax, sMax int

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num > fMax {
			fMax, sMax = num, fMax
		} else if num > sMax {
			sMax = num
		}
	}

	return (fMax - 1) * (sMax - 1)
}
