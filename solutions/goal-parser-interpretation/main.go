package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(interpret("G()(al)"))
	fmt.Println(interpret("G()()()()(al)"))
	fmt.Println(interpret("(al)G(al)()()G"))
}

func interpret(command string) string {
	var result strings.Builder
	var counter int

	for counter < len(command) {
		charByte := command[counter]
		if charByte == 'G' {
			result.WriteByte('G')
			counter++
		} else if charByte == '(' && command[counter+1] == ')' {
			result.WriteByte('o')
			counter += 2
		} else {
			result.Write([]byte{'a', 'l'})
			counter += 4
		}
	}

	return result.String()
}
