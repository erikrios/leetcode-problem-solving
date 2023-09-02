package main

import "fmt"

func main() {
	fmt.Println(isAcronym([]string{"alice", "bob", "charlie"}, "abc"))
	fmt.Println(isAcronym([]string{"an", "apple"}, "a"))
	fmt.Println(isAcronym([]string{"never", "gonna", "give", "up", "on", "you"}, "ngguoy"))
}

func isAcronym(words []string, s string) bool {
	n := len(words)
	m := len(s)

	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		if s[i] != words[i][0] {
			return false
		}
	}

	return true
}
