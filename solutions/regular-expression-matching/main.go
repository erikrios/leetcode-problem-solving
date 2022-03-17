package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
	fmt.Println(isMatch("ab", "*a"))
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := !(len(s) == 0) && (p[0] == s[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, string(p[2:])) || (firstMatch && isMatch(string(s[1:]), p))
	} else {
		return firstMatch && isMatch(string(s[1:]), string(p[1:]))
	}
}
