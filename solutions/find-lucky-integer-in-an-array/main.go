package main

import "fmt"

func main() {
	fmt.Println(findLucky([]int{2, 2, 3, 4}))
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))
}

func findLucky(arr []int) int {
	freqs := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		freqs[arr[i]]++
	}

	luckyNum := -1
	for k, v := range freqs {
		if k != v {
			continue
		}

		if v > luckyNum {
			luckyNum = k
		}
	}

	return luckyNum
}
