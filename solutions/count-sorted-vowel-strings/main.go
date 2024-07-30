package main

import "fmt"

func main() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(3))
	fmt.Println(countVowelStrings(4))
	fmt.Println(countVowelStrings(33))
}

const vowels = 5

func countVowelStrings(n int) int {
	return count(n, 0, 0)
}

func count(n, i, total int) int {
	if total == n {
		return 1
	}

	total++

	var sum int
	for j := i; j < vowels; j++ {
		sum += count(n, j, total)
	}

	return sum
}
