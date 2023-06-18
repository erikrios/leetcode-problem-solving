package main

import "fmt"

func main() {
	fmt.Println(maximumValue([]string{"alic3", "bob", "3", "4", "00000"}))
	fmt.Println(maximumValue([]string{"1", "01", "001", "0001"}))
}

func maximumValue(strs []string) int {
	var max int

	for i := 0; i < len(strs); i++ {
		str := strs[i]
		var val int

		if isNum(str) {
			val = convToDecimal(str)
		} else {
			val = len(str)
		}

		if val > max {
			max = val
		}
	}

	return max
}

func isNum(str string) bool {
	for i := 0; i < len(str); i++ {
		if !(str[i] >= '0' && str[i] <= '9') {
			return false
		}
	}

	return true
}

func convToDecimal(str string) int {
	var res int
	multiplier := 1
	const base = 10

	for i := 0; i < len(str); i++ {
		char := str[len(str)-1-i]
		num := int(char-'0') * multiplier
		res += num
		multiplier *= base
	}

	return res
}
