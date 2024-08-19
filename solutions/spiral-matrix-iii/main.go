package main

import "fmt"

func main() {
	fmt.Println(spiralMatrixIII(1, 4, 0, 0))
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}

type Direction int

const (
	Right Direction = iota
	Bottom
	Left
	Top
)

var directions = [4]Direction{Right, Bottom, Left, Top}

func spiralMatrixIII(rows, cols, rStart, cStart int) [][]int {
	results := make([][]int, rows*cols)

	visitCoordinates(
		rows,
		cols,
		rStart,
		cStart,
		0,
		0,
		2,
		0,
		results,
	)

	return results
}

func visitCoordinates(
	rows,
	cols,
	rStart,
	cStart,
	resIndex,
	curStep,
	totalSteps,
	dirIndex int,
	results [][]int,
) {
	if resIndex == (rows * cols) {
		return
	}

	if isInside(rows, cols, rStart, cStart) {
		results[resIndex] = []int{rStart, cStart}
		resIndex++
	}

	if curStep == totalSteps/2 {
		dirIndex++
	}

  if curStep == totalSteps {
		dirIndex++
		totalSteps += 2
    curStep = 1
	} else {
    curStep++
  }

	switch directions[dirIndex%len(directions)] {
	case Right:
		cStart++
	case Bottom:
		rStart++
	case Left:
		cStart--
	case Top:
		rStart--
	}

	visitCoordinates(
		rows,
		cols,
		rStart,
		cStart,
		resIndex,
		curStep,
		totalSteps,
		dirIndex,
		results,
	)
}

func isInside(rows, cols, i, j int) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols
}
