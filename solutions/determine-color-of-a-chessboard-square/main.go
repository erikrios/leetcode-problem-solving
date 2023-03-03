package main

import "fmt"

func main() {
	fmt.Println(squareIsWhite("a1"))
	fmt.Println(squareIsWhite("h3"))
	fmt.Println(squareIsWhite("c7"))
}

func squareIsWhite(coordinates string) bool {
	a, b := int(coordinates[0]), int(coordinates[1])
	res := a - b
	if res < 0 {
		res *= -1
	}

	return res%2 != 0
}
