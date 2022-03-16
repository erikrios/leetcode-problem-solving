package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("    -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("   +0 123"))
}

func myAtoi(s string) int {
	s = clean(s)

	if len(s) == 0 {
		return 0
	}

	var isNegative bool

	if s[0] == '-' {
		isNegative = true
	} else {
		isNegative = false
	}

	if s[0] == '-' || s[0] == '+' {
		s = removeLeadingZero(string(s[1:]))
	} else {
		s = removeLeadingZero(s)
	}

	res, _ := strconv.Atoi(s)

	if isNegative {
		res *= -1
		if res < math.MinInt32 {
			return math.MinInt32
		}
		return res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func clean(s string) string {
	res := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && len(res) == 0 {
			continue
		} else if s[i] >= '0' && s[i] <= '9' {
			res = append(res, s[i])
		} else {
			if (s[i] == '-' || s[i] == '+') && len(res) == 0 {
				res = append(res, s[i])
			} else {
				break
			}
		}
	}

	return string(res)
}

func removeLeadingZero(s string) string {
	var counter uint8

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			counter++
		} else {
			break
		}
	}

	return string(s[counter:])
}
