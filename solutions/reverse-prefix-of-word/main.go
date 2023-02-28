package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}

func reversePrefix(word string, ch byte) string {
	var isShouldReverse bool
	var incIndex int

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			isShouldReverse = true
			incIndex = i
			break
		}
	}

	if !isShouldReverse {
		return word
	}

	sb := strings.Builder{}
	for i := incIndex; i >= 0; i-- {
		sb.WriteByte(word[i])
	}

	for i := incIndex + 1; i < len(word); i++ {
		sb.WriteByte(word[i])
	}

	return sb.String()
}
