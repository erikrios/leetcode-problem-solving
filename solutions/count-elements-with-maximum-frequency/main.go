package main

import "fmt"

func main() {
	fmt.Println(maxFrequencyElements([]int{1, 2, 2, 3, 1, 4}))
	fmt.Println(maxFrequencyElements([]int{1, 2, 3, 4, 5}))
}

func maxFrequencyElements(nums []int) int {
	n := len(nums)
	freqs := make(map[int]int)
	var maxFreq int

	for i := 0; i < n; i++ {
		num := nums[i]
		freqs[num]++
		freq := freqs[num]
		if maxFreq < freq {
			maxFreq = freq
		}
	}

	var total int
	for _, v := range freqs {
		if v == maxFreq {
			total += v
		}
	}

	return total
}
