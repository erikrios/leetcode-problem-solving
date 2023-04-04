package main

import "fmt"

func main() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
	fmt.Println(twoOutOfThree([]int{3, 1}, []int{2, 3}, []int{1, 2}))
	fmt.Println(twoOutOfThree([]int{1, 2, 2}, []int{4, 3, 3}, []int{5}))
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	const length = 101
	counters := make([][3]bool, length, length)
	lenNum1, lenNum2, lenNum3 := len(nums1), len(nums2), len(nums3)

	var i int
	for {
		if i >= lenNum1 && i >= lenNum2 && i >= lenNum3 {
			break
		}

		if i < lenNum1 {
			num1 := nums1[i]
			counters[num1][0] = true
		}

		if i < lenNum2 {
			num2 := nums2[i]
			counters[num2][1] = true
		}

		if i < lenNum3 {
			num3 := nums3[i]
			counters[num3][2] = true
		}

		i++
	}

	results := make([]int, 0)
	for i := 0; i < length; i++ {
		flags := counters[i]
		var total int

		if flags[0] {
			total++
		}

		if flags[1] {
			total++
		}

		if flags[2] {
			total++
		}

		if total >= 2 {
			results = append(results, i)
		}
	}

	return results
}
