package main

import "fmt"

func main() {
	fmt.Println(scoreOfString("hello"))
	fmt.Println(scoreOfString("zaz"))
}

func scoreOfString(s string) int {
	var sum int

	for i := 0; i < len(s)-1; i++ {
		cur := s[i]
		next := s[i+1]

		var diff int
		if cur > next {
			diff = int(cur - next)
		} else {
			diff = int(next - cur)
		}

		sum += diff
	}

	return sum
}
