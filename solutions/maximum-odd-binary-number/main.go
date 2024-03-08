package main

import "fmt"

func main() {
	fmt.Println(maximumOddBinaryNumber("101"))
	fmt.Println(maximumOddBinaryNumber("0101"))
}

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	var bitSetCount int

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			bitSetCount++
		}
	}

	res := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		if bitSetCount > 1 || i == n-1 {
			res = append(res, '1')
			bitSetCount--
		} else {
			res = append(res, '0')
		}
	}

	return string(res)
}
