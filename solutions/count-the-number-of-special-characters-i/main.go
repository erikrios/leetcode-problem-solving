package main

import "fmt"

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
	fmt.Println(numberOfSpecialChars("abc"))
	fmt.Println(numberOfSpecialChars("abBCab"))
}

func numberOfSpecialChars(word string) int {
	isExistsMapper := make(map[byte]struct{})
	upperCaseMapper := make(map[byte]struct{})
	lowerCaseMapper := make(map[byte]struct{})

	n := len(word)

	for i := 0; i < n; i++ {
		char := word[i]
		isLower := isLower(char)
		lowerChar := char
		if !isLower {
			lowerChar = toLower(char)
		}

		if _, ok := isExistsMapper[lowerChar]; ok {
			continue
		}

		if isLower {
			lowerCaseMapper[lowerChar] = struct{}{}
			if _, ok := upperCaseMapper[lowerChar]; !ok {
				continue
			}
		} else {
			upperCaseMapper[lowerChar] = struct{}{}
			if _, ok := lowerCaseMapper[lowerChar]; !ok {
				continue
			}
		}

		isExistsMapper[lowerChar] = struct{}{}
	}

	return len(isExistsMapper)
}

func isLower(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func toLower(char byte) byte {
	return char - 'A' + 'a'
}
