package main

import "fmt"

func main() {
	fmt.Println(isPrefixOfWord("i love eating burger", "burg"))
	fmt.Println(isPrefixOfWord("this problem is an easy problem", "pro"))
	fmt.Println(isPrefixOfWord("i am tired", "you"))
}

func isPrefixOfWord(sentence, searchWord string) int {
	var shouldSkip bool
	var j int
	n := len(sentence)
	var idx int

	for i := 0; i < n; i++ {
		char := sentence[i]
		if i > 0 && char == ' ' {
			shouldSkip = false
			idx++
			j = 0
			continue
		}

		if shouldSkip {
			continue
		}

		if char != searchWord[j] {
			j = 0
			shouldSkip = true
			continue
		}

		if i == n-1 && j == len(searchWord)-1 {
			return idx + 1
		} else if i < n-1 && sentence[i+1] == ' ' && j == len(searchWord)-1 {
			return idx + 1
		} else if i < n-1 && j == len(searchWord)-1 {
			return idx + 1
		}

		j++
	}

	return -1
}
