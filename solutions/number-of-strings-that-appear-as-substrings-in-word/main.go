package main

import "fmt"

func main() {
	fmt.Println(numOfStrings([]string{"a", "abc", "bc", "d"}, "abc"))
	fmt.Println(numOfStrings([]string{"a", "b", "c"}, "aaaaabbbbb"))
	fmt.Println(numOfStrings([]string{"a", "a", "a"}, "ab"))
}

func numOfStrings(patterns []string, word string) int {
	mapper := make(map[string]bool)

	for i := 0; i < len(word); i++ {
		for j := 0; j+i+1 <= len(word); j++ {
			mapper[string(word[j:j+i+1])] = true
		}
	}

	var counter int
	for i := 0; i < len(patterns); i++ {
		subStr := patterns[i]
		if _, ok := mapper[subStr]; ok {
			counter++
		}
	}

	return counter
}
