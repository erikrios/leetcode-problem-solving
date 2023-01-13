package main

import (
	"fmt"
)

func main() {
	fmt.Println(cellsInRange("K1:L2"))
	fmt.Println(cellsInRange("A1:F1"))
}

func cellsInRange(s string) []string {
	cap := (s[3] - s[0] + 1) * (s[4] - s[1] + 1)
	results := make([]string, cap, cap)

	var counter int
	for i := s[0]; i <= s[3]; i++ {
		for j := s[1]; j <= s[4]; j++ {
			results[counter] = string(i) + string(j)
			counter++
		}
	}

	return results
}
