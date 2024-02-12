package main

import "fmt"

func main() {
	fmt.Println(findKOr([]int{7, 12, 9, 8, 9, 15}, 4))
	fmt.Println(findKOr([]int{2, 12, 1, 11, 4, 5}, 6))
	fmt.Println(findKOr([]int{10, 8, 5, 9, 11, 6, 8}, 1))
}

func findKOr(nums []int, k int) int {
	n := len(nums)
	numsBits := make([][]int, n)
	var maxLen int

	for i := 0; i < n; i++ {
		num := nums[i]
		numsBits[i] = getBits(num, k)
		length := len(numsBits[i])
		if maxLen < length {
			maxLen = length
		}
	}

	var kOr int
	for i := 0; i < maxLen; i++ {
		var counter int
		for j := 0; j < len(numsBits); j++ {
			numBits := numsBits[j]
			if i < len(numBits) && numBits[i] == 1 {
				counter++
			}
		}

		if counter >= k {
			val := 1
			for l := 0; l < i; l++ {
				val *= 2
			}
			kOr += val
		}
	}

	return kOr
}

func getBits(num int, length int) []int {
	bits := []int{}
	for num > 0 {
		digit := num % 2
		num /= 2
		bits = append(bits, digit)
	}

	return bits
}
