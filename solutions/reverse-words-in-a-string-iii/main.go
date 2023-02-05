package main

import "fmt"

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))
}

func reverseWords(s string) string {
	spacePtrs := make([]int, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			spacePtrs = append(spacePtrs, i)
		}
	}

	spacePtrs = append(spacePtrs, len(s))

	var spacePtr int
	ptr := spacePtrs[spacePtr] - 1
	results := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			results[ptr] = s[i]
			ptr--
		} else {
			results[i] = ' '
			spacePtr++
			ptr = spacePtrs[spacePtr] - 1
		}
	}

	return string(results)
}
