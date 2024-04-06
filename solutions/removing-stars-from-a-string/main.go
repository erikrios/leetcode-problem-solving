package main

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}

func removeStars(s string) string {
	results := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		switch char := s[i]; char {
		case '*':
			results = results[:len(results)-1]
		default:
			results = append(results, char)
		}
	}

	return string(results)
}
