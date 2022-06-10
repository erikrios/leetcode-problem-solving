package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	var length int
	startingEmpty := s[len(s)-1] == ' '

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
			startingEmpty = false
		}

		if s[i] == ' ' && startingEmpty == false {
			break
		}
	}

	return length
}
