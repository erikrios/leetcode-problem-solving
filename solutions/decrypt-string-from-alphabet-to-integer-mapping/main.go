package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
}

func freqAlphabets(s string) string {
	var builder strings.Builder

	var i int

	for i < len(s) {
		var vNum int

		if i+2 < len(s) && s[i+2] == '#' {
			vNum, _ = strconv.Atoi(string(s[i]) + string(s[i+1]))
			i += 3
		} else {
			vNum, _ = strconv.Atoi(string(s[i]))
			i++
		}

		v := 'a' + byte(vNum-1)
		builder.WriteByte(v)
	}

	return builder.String()
}
