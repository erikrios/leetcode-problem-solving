package main

import "fmt"

func main() {
	fmt.Println(countGoodRectangles([][]int{
		{5, 8},
		{3, 9},
		{5, 12},
		{16, 5},
	}))

	fmt.Println(countGoodRectangles([][]int{
		{2, 3},
		{3, 7},
		{4, 3},
		{3, 7},
	}))
}

func countGoodRectangles(rectangles [][]int) int {
	var counter int
	var max int
	for i := 0; i < len(rectangles); i++ {
		rect := rectangles[i]

		var square int
		if rect[0] < rect[1] {
			square = rect[0]
		} else {
			square = rect[1]
		}

		if square > max {
			max = square
			counter = 1
		} else if square == max {
			counter++
		}
	}

	return counter
}
