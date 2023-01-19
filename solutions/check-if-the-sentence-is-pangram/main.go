package main

import "fmt"

func main() {
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(checkIfPangram("leetcode"))
}

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}

	alphabet := make(map[rune]bool)

	for i := 'a'; i <= 'z'; i++ {
		alphabet[i] = false
	}

	for _, char := range sentence {
		alphabet[char] = true
	}

	result := true

	for _, v := range alphabet {
		if !v {
			result = false
			break
		}
	}

	return result
}
