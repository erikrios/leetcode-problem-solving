package main

import "fmt"

func main() {
	fmt.Println(executeInstructions(3, []int{0, 1}, "RRDDLU"))
	fmt.Println(executeInstructions(2, []int{1, 1}, "LURD"))
	fmt.Println(executeInstructions(1, []int{0, 0}, "LRUD"))
}

func executeInstructions(n int, startPos []int, s string) []int {
	totalMovs := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		totalMove := totalMovement(n, startPos, []byte(s[i:]))
		totalMovs[i] = totalMove
	}

	return totalMovs
}

func totalMovement(n int, startPos []int, instructions []byte) int {
	var counter int
	startRow, startCol := startPos[0], startPos[1]

	for i := 0; i < len(instructions); i++ {
		switch instruction := instructions[i]; instruction {
		case 'L':
			startCol--
		case 'R':
			startCol++
		case 'U':
			startRow--
		default:
			startRow++
		}

		if startRow >= n || startCol >= n || startRow < 0 || startCol < 0 {
			break
		}

		counter++
	}

	return counter
}
