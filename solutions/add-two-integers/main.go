package main

import "fmt"

func main() {
	fmt.Println(sum(12, 5))
	fmt.Println(sum(-10, 4))
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
