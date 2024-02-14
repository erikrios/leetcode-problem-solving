package main

import "fmt"

func main() {
	fmt.Println(maxLengthBetweenEqualCharacters("aa"))
	fmt.Println(maxLengthBetweenEqualCharacters("abca"))
	fmt.Println(maxLengthBetweenEqualCharacters("cbzxy"))
}

func maxLengthBetweenEqualCharacters(s string) int {
	n := len(s)
	chars := make(map[byte][]int)

	for i := 0; i < n; i++ {
		char := s[i]
		chars[char] = append(chars[char], i)
	}

	max := -1
	for _, indices := range chars {
		m := len(indices)
		if m > 1 {
			ranges := indices[m-1] - indices[0] - 1
			if ranges > max {
				max = ranges
			}
		}
	}

	return max
}
