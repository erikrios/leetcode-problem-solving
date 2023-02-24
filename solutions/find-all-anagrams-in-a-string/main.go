package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("baa", "aa"))
}

func findAnagrams(s string, p string) []int {
	results := make([]int, 0)

	for i := 0; i < len(s)-len(p)+1; i++ {
		mapper := make(map[byte]uint16)

		for j := i; j < i+len(p); j++ {
			mapper[s[j]] += 1
		}

		isAnagram := true
		for j := 0; j < len(p); j++ {
			if _, ok := mapper[p[j]]; !ok {
				isAnagram = false
				break
			}
			if mapper[p[j]] -= 1; mapper[p[j]] <= 0 {
				delete(mapper, p[j])
			}
		}

		if isAnagram {
			results = append(results, i)
		}
	}

	return results
}
