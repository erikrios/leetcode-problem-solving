package main

import "fmt"

func main() {
	subRect := Constructor([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	fmt.Println(subRect)
	fmt.Println(subRect.GetValue(0, 2))
	subRect.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Println(subRect)
	fmt.Println(subRect.GetValue(0, 2))
	fmt.Println(subRect.GetValue(3, 1))
	subRect.UpdateSubrectangle(3, 0, 3, 2, 10)
	fmt.Println(subRect)
	fmt.Println(subRect.GetValue(3, 1))
	fmt.Println(subRect.GetValue(0, 2))

	fmt.Println()

	subRect2 := Constructor([][]int{{3, 9, 4}, {5, 6, 10}})
	fmt.Println(subRect2)
	subRect2.UpdateSubrectangle(1, 1, 1, 1, 5)
	fmt.Println(subRect2)
	fmt.Println(subRect2.GetValue(1, 0))
	fmt.Println(subRect2.GetValue(1, 0))
	subRect2.UpdateSubrectangle(0, 0, 1, 0, 6)
	fmt.Println(subRect2)
	fmt.Println(subRect2.GetValue(1, 0))
	fmt.Println(subRect2.GetValue(1, 0))
}

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}
