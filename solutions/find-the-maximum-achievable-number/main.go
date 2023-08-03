package main

import "fmt"

func main() {
	fmt.Println(theMaximumAchievableX(4, 1))
	fmt.Println(theMaximumAchievableX(3, 2))
}

func theMaximumAchievableX(num int, t int) int {
	return num + 2*t
}
