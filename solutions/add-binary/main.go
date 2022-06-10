package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	sb := strings.Builder{}

	i, j, carry := len(a)-1, len(b)-1, 0

	for i >= 0 || j >= 0 {
		sum := carry
		if j >= 0 {
			sum += int(b[j]) - '0'
			j--
		}
		if i >= 0 {
			sum += int(a[i]) - '0'
			i--
		}
		sb.WriteString(strconv.Itoa(sum % 2))
		carry = sum / 2
	}

	if carry != 0 {
		sb.WriteString(strconv.Itoa(carry))
	}
	str := sb.String()

	var res string

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	return res
}
