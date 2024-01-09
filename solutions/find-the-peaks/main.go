package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPeaks([]int{2, 4, 4}))
	fmt.Println(findPeaks([]int{1, 4, 3, 8, 5}))
	fmt.Println(findPeaks([]int{3, 3, 4, 2}))
}

func findPeaks(mountain []int) []int {
	n := len(mountain)
	capacity := math.Round(float64(n) / float64(3))
	peaks := make([]int, 0, int(capacity))

	for i := 1; i < n-1; {
		cur := mountain[i]
		prev := mountain[i-1]
		next := mountain[i+1]

		if cur > prev && cur > next {
			peaks = append(peaks, i)
			i += 2
		} else if cur == next {
			i += 2
		} else {
			i++
		}
	}

	return peaks
}
