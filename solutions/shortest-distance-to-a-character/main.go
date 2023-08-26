package main

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
	fmt.Println(shortestToChar("aaab", 'b'))
}

func shortestToChar(s string, c byte) []int {
	cPositions := make([]int, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] == c {
			cPositions = append(cPositions, i)
		}
	}

	results := make([]int, 0, n)
	for i := 0; i < n; i++ {

		min := n
		for j := 0; j < len(cPositions); j++ {
			cPos := cPositions[j]
			distance := cPos - i
			if distance < 0 {
				distance *= -1
			}

			if distance < min {
				min = distance
			}
		}

		results = append(results, min)

	}

	return results
}
