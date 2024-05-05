package main

import "fmt"

func main() {
	fmt.Println(reverseOnlyLetters("ab-cd"))
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
	fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}

func reverseOnlyLetters(s string) string {
	n := len(s)
	reversedStr := []byte(s)

	start, end := 0, n-1

	for start < end {
		startChar, endChar := reversedStr[start], reversedStr[end]

		isStartCharLetter, isEndCharLetter := isLetter(startChar), isLetter(endChar)

		if !isStartCharLetter && !isEndCharLetter {
			start++
			end--
		} else if isStartCharLetter && !isEndCharLetter {
			end--
		} else if !isStartCharLetter && isEndCharLetter {
			start++
		} else {
			reversedStr[start], reversedStr[end] = endChar, startChar
			start++
			end--
		}
	}

	return string(reversedStr)
}

func isLetter(char byte) bool {
	return (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z')
}
