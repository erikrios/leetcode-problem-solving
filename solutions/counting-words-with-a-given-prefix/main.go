package main

import "fmt"

func main() {
	fmt.Println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
	fmt.Println(prefixCount([]string{"leetcode", "win", "loops", "success"}, "code"))
}

func prefixCount(words []string, pref string) int {
	var counter int

	for i := 0; i < len(words); i++ {
		word := words[i]
		if len(word) < len(pref) {
			continue
		}

		hasPref := true
		for j := 0; j < len(pref); j++ {
			if word[j] != pref[j] {
				hasPref = false
				break
			}
		}

		if hasPref {
			counter++
		}
	}

	return counter
}
