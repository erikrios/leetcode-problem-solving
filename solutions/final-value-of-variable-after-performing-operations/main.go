package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
	fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X++"}))
	fmt.Println(finalValueAfterOperations([]string{"X++", "++X", "--X", "X--"}))
}

func finalValueAfterOperations(operations []string) int {
	var sum int

	for _, operation := range operations {
		if strings.Contains(operation, "+") {
			sum++
		} else {
			sum--
		}
	}

	return sum
}
