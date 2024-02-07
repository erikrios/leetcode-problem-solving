package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(frequencySort("tree"))
	fmt.Println(frequencySort("cccaaa"))
	fmt.Println(frequencySort("Aabb"))
}

func frequencySort(s string) string {
	charMapper := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charMapper[s[i]]++
	}

	results := make(
		[]struct {
			char  byte
			count int
		},
		0,
		len(charMapper),
	)

	var n int
	for k, v := range charMapper {
		n += v
		results = append(results, struct {
			char  byte
			count int
		}{k, v})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})

	freqRes := make([]byte, 0, n)
	for i := 0; i < len(results); i++ {
		result := results[i]
		for j := 0; j < result.count; j++ {
			freqRes = append(freqRes, result.char)
		}
	}

	return string(freqRes)
}
