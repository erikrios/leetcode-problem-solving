package main

import "fmt"

func main() {
	fmt.Println(checkString("aaabbb"))
	fmt.Println(checkString("abab"))
	fmt.Println(checkString("bbb"))
}

func checkString(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		cur := s[i]
		next := s[i+1]
		if cur > next {
			return false
		}
	}

	return true
}
