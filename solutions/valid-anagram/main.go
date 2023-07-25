package main

import "fmt"

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapper := make(map[byte][2]int)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		vals := mapper[sChar]
		vals[0]++
		mapper[sChar] = vals

		vals = mapper[tChar]
		vals[1]++
		mapper[tChar] = vals
	}

	for _, vals := range mapper {
		if vals[0] != vals[1] {
			return false
		}
	}

	return true
}
