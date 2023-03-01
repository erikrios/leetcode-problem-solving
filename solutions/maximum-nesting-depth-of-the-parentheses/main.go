package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
}

func maxDepth(s string) int {
	var counter int
	var max int

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char == '(' {
			counter++
			if counter > max {
				max = counter
			}
		} else if char == ')' {
			counter--
		}
	}

	return max
}
