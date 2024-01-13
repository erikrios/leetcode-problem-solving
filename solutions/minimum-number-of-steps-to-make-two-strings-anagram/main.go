package main

import "fmt"

func main() {
	fmt.Println(minSteps("bab", "aba"))
	fmt.Println(minSteps("leetcode", "practice"))
	fmt.Println(minSteps("anagram", "mangaar"))
}

func minSteps(s, t string) int {
	n := len(s)

	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < n; i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	var steps int
	for k, v := range sMap {
		if val := tMap[k]; v-val > 0 {
			num := v - val
			steps += num
		}
	}

	return steps
}
