package main

import "fmt"

func main() {
	fmt.Println(maximumUnits([][]int{{1, 3}, {2, 2}, {3, 1}}, 4))
	fmt.Println(maximumUnits([][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10))
	fmt.Println(maximumUnits([][]int{{2, 1}, {4, 4}, {3, 1}, {4, 1}, {2, 4}, {3, 4}, {1, 3}, {4, 3}, {5, 3}, {5, 3}}, 13))
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	counts := make([]int, 1001, 1001)

	for i := 0; i < len(boxTypes); i++ {
		boxType := boxTypes[i]
		counts[boxType[1]] += boxType[0]
	}

	var max int
	var boxesCounter int
	for i := 1000; i >= 0; i-- {
		if boxesCounter == truckSize {
			return max
		}

		totalBoxes := counts[i]
		unitAmountPerBoxes := i

		if totalBoxes == 0 {
			continue
		}

		if boxesCounter == truckSize {
			return max
		}

		if boxesCounter+totalBoxes <= truckSize {
			max += totalBoxes * unitAmountPerBoxes
			boxesCounter += totalBoxes
		} else {
			remainingBoxes := truckSize - boxesCounter
			max += remainingBoxes * unitAmountPerBoxes
			boxesCounter += remainingBoxes
		}
	}

	return max
}
