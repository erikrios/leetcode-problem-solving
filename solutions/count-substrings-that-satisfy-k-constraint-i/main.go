package main

import "fmt"

func main() {
	fmt.Println(countKConstraintSubstrings("10101", 1))
	fmt.Println(countKConstraintSubstrings("1010101", 2))
	fmt.Println(countKConstraintSubstrings("11111", 1))
}

func countKConstraintSubstrings(s string, k int) int {
	windowStart := 0

	var zerosCount, onesCount, count int
	for windowEnd := 0; windowEnd < len(s); windowEnd++ {
		ch := s[windowEnd]
		if ch == '0' {
			zerosCount++
		} else {
			onesCount++
		}

		for zerosCount > k && onesCount > k {
			if s[windowStart] == '0' {
				zerosCount--
			} else {
				onesCount--
			}
			windowStart++
		}

		count += (windowEnd - windowStart + 1)
	}

	return count
}
