package main

import "fmt"

func main() {
	fmt.Println(countAsteriks("l|*e*et|c**o|*de|"))
	fmt.Println(countAsteriks("iamprogrammer"))
	fmt.Println(countAsteriks("yo|uar|e**|b|e***au|tifu|l"))
}

func countAsteriks(s string) int {
	var isBetween bool
	var counter int

	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			isBetween = !isBetween
		}

		if !isBetween && s[i] == '*' {
			counter++
		}
	}

	return counter
}
