package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}

func largestAltitude(gain []int) int {
	var highest, alt int

	for i := 0; i < len(gain); i++ {
		alt += gain[i]
		if alt > highest {
			highest = alt
		}
	}

	return highest
}
