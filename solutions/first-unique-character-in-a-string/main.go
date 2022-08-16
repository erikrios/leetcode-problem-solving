package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	mapper := make(map[byte]uint16)

	for i := 0; i < len(s); i++ {
		mapper[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if count := mapper[s[i]]; count == 1 {
			return i
		}
	}

	return -1
}
