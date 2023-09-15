package main

import "fmt"

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}

func divideString(s string, k int, fill byte) []string {
	n := len(s)
	m := n / k
	div := n % k
	rem := k - div
	if div != 0 {
		m++
	}

	res := make([]string, 0, m)

	chars := make([]byte, 0, k)
	for i := 0; i < n; i++ {
		chars = append(chars, s[i])
		if (i+1)%k == 0 {
			res = append(res, string(chars))
			chars = make([]byte, 0, k)
		}
	}

	if div != 0 {
		chars = fillByte(chars, fill, rem)
		res = append(res, string(chars))
	}

	return res
}

func fillByte(chars []byte, fill byte, times int) []byte {
	for i := 0; i < times; i++ {
		chars = append(chars, fill)
	}

	return chars
}
