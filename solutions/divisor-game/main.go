package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
}

func divisorGame(n int) bool {
	return n%2 == 0
}
