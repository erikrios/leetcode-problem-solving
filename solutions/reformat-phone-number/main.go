package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reformatNumber("1-23-45 6"))
	fmt.Println(reformatNumber("123 4-567"))
	fmt.Println(reformatNumber("123 4-5678"))
}

func reformatNumber(number string) string {
	n := len(number)
	nums := make([]byte, 0, n)

	for i := 0; i < n; i++ {
		char := number[i]
		if char == ' ' || char == '-' {
			continue
		}

		nums = append(nums, char)
	}

	n = len(nums)
	var builder strings.Builder
	dashPos := 3
	rem := n % dashPos

	for i := 0; i < n; i++ {
		char := nums[i]
		builder.WriteByte(char)
		remLen := n - i - 1

		if rem == 1 && remLen < 4 && i < n-1 {
			if remLen == 2 {
				builder.WriteByte('-')
			}
			continue
		}

		if (i+1)%dashPos == 0 && i < n-1 {
			builder.WriteByte('-')
		}
	}

	return builder.String()
}
