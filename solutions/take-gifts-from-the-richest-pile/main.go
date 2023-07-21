package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(pickGifts([]int{25, 64, 9, 4, 100}, 4))
	fmt.Println(pickGifts([]int{1, 1, 1, 1}, 4))
}

func pickGifts(gifts []int, k int) int64 {
	sort.Ints(gifts)

	for i := 0; i < k; i++ {
		gift := gifts[len(gifts)-1]
		sqrt := int(math.Sqrt(float64(gift)))
		gifts[len(gifts)-1] = sqrt
		shift(gifts)
	}

	var sum int64

	for i := 0; i < len(gifts); i++ {
		sum += int64(gifts[i])
	}

	return sum
}

func shift(nums []int) {
	for i := len(nums) - 1; i >= 1; i-- {
		cur := nums[i]
		prev := nums[i-1]
		if cur < prev {
			nums[i], nums[i-1] = prev, cur
		} else {
			break
		}
	}
}
