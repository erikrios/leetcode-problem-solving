package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("aa", "a"))
	fmt.Println(minWindow("aa", "aa"))
}

func minWindow(s, t string) string {
	m := len(s)
	n := len(t)

	if n == 0 {
		return ""
	}

	countT, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < n; i++ {
		countT[t[i]]++
	}

	have, need := 0, len(countT)

	res, resLen := [2]int{-1, -1}, math.MaxInt
	l := 0

	for r := 0; r < m; r++ {
		c := s[r]
		window[c]++

		if v, ok := countT[c]; ok && window[c] == v {
			have++
		}

		for have == need {
			if r-l+1 < resLen {
				res = [2]int{l, r}
				resLen = r - l + 1
			}
			window[s[l]]--
			if v, ok := countT[s[l]]; ok && window[s[l]] < v {
				have--
			}
			l += 1
		}
	}

	l, r := res[0], res[1]
	if resLen == math.MaxInt {
		return ""
	}

	return string(s[l : r+1])
}
