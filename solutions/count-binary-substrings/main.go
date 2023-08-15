package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("00110"))
}

func countBinarySubstrings(s string) int {
	cur, pre, res := 1, 0, 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			min := cur
			if pre < cur {
				min = pre
			}
			res += min
			pre = cur
			cur = 1
		}
	}

	min := cur
	if pre < cur {
		min = pre
	}

	return res + min
}
