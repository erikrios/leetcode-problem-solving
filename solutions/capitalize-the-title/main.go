package main

import "fmt"

func main() {
	fmt.Println(capitalizeTitle("capiTalIze tHe titLe"))
	fmt.Println(capitalizeTitle("First Letter of Each Word"))
	fmt.Println(capitalizeTitle("i lOve leetcode"))
	fmt.Println(capitalizeTitle("L hV"))
}

func capitalizeTitle(title string) string {
	n := len(title)
	results := []byte(title)

	var startIdx, endIdx int
	for i := 0; i < n; i++ {
		char := title[i]

		if char == ' ' || i == n-1 {
			endIdx = i - 1
			if i == n-1 {
				endIdx = i
				if !isLower(char) {
					results[i] = toLower(char)
				}
			}

			length := endIdx - startIdx + 1
			if length > 2 {
				if !isLower(results[startIdx]) {
					continue
				}
				results[startIdx] = toUpper(results[startIdx])
			}
			startIdx = i + 1

			continue
		}

		if isLower(char) {
			continue
		}
		results[i] = toLower(char)
	}

	return string(results)
}

func isLower(char byte) bool {
	return char >= 'a' && char <= 'z'
}

func toUpper(char byte) byte {
	return char - 'a' + 'A'
}

func toLower(char byte) byte {
	return char - 'A' + 'a'
}
