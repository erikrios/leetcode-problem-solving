package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints("B0R0G0R9R0B0G0"))
	fmt.Println(countPoints("G4"))
}

func countPoints(rings string) int {
	rods := make([][]byte, 10, 10)

	for i := 0; i < len(rings); i += 2 {
		rodPos := rings[i+1] - '0'
		rods[rodPos] = append(rods[rodPos], rings[i])
	}

	var counter int

	for i := 0; i < len(rods); i++ {
		rod := rods[i]

		var red, green, blue int
		for j := 0; j < len(rod); j++ {
			switch rod[j] {
			case 'R':
				red++
			case 'G':
				green++
			default:
				blue++
			}
		}

		if red > 0 && green > 0 && blue > 0 {
			counter++
		}
	}

	return counter
}
