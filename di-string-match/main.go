package main

import (
	"fmt"
)

func main() {
	fmt.Println(diStringMatch("IDID"))
	fmt.Println(diStringMatch("III"))
	fmt.Println(diStringMatch("DDI"))
	fmt.Println(diStringMatch("DDD"))
}

func diStringMatch(s string) []int {
	n := len(s)
	left, right := 0, n
	res := make([]int, n+1)

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
	}

	res[n] = left
	return res
}
