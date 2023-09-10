package main

import "fmt"

func main() {
	fmt.Println(garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3}))
	fmt.Println(garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10}))
}

func garbageCollection(garbage []string, travel []int) int {
	var totalPickTimeP int
	var totalMoveTimeP int
	var tempMoveTimeP int

	var totalPickTimeG int
	var totalMoveTimeG int
	var tempMoveTimeG int

	var totalPickTimeM int
	var totalMoveTimeM int
	var tempMoveTimeM int

	for i := 0; i < len(garbage); i++ {
		house := garbage[i]

		var totalCollected int
		garbage[i], totalCollected = collects(house, 'P')
		totalPickTimeP += totalCollected

		if i > 0 {
			tempMoveTimeP += travel[i-1]
		}

		if totalCollected > 0 {
			totalMoveTimeP += tempMoveTimeP
			tempMoveTimeP = 0
		}

		totalCollected = 0
		garbage[i], totalCollected = collects(house, 'G')
		totalPickTimeG += totalCollected

		if i > 0 {
			tempMoveTimeG += travel[i-1]
		}

		if totalCollected > 0 {
			totalMoveTimeG += tempMoveTimeG
			tempMoveTimeG = 0
		}

		totalCollected = 0
		garbage[i], totalCollected = collects(house, 'M')
		totalPickTimeM += totalCollected

		if i > 0 {
			tempMoveTimeM += travel[i-1]
		}

		if totalCollected > 0 {
			totalMoveTimeM += tempMoveTimeM
			tempMoveTimeM = 0
		}
	}

	totalP := totalPickTimeP + totalMoveTimeP
	totalG := totalPickTimeG + totalMoveTimeG
	totalM := totalPickTimeM + totalMoveTimeM

	return totalP + totalG + totalM
}

func collects(
	garbage string,
	truckType byte,
) (
	newGarbage string,
	totalCollected int,
) {
	gbs := make([]byte, 0)

	for i := 0; i < len(garbage); i++ {
		if gb := garbage[i]; gb != truckType {
			gbs = append(gbs, gb)
		}
	}

	newGarbage = string(gbs)
	totalCollected = len(garbage) - len(newGarbage)

	return
}
