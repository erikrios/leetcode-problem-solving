package main

import "fmt"

func main() {
	fmt.Println(greatestLetter("lEeTcOdE"))
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}

func greatestLetter(s string) string {
	n := len(s)
	mapper := make(map[byte]struct{}, n)

	for i := 0; i < n; i++ {
		char := s[i]
		mapper[char] = struct{}{}
	}

	var res byte
	for k, _ := range mapper {
		if k >= 'A' && k <= 'Z' {
			lower := k - 'A' + 'a'
			if _, ok := mapper[lower]; ok {
				if k > res {
					res = k
				}
			}
		}
	}

	if res == 0 {
		return ""
	}

	return string(res)
}
