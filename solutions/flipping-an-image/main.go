package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}))
}

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		var jConstraint int
		if len(image)%2 == 0 {
			jConstraint = len(image[i]) / 2
		} else {
			jConstraint = len(image[i])/2 + 1
		}

		for j := 0; j < jConstraint; j++ {
			front := image[i][j]
			back := image[i][len(image[i])-1-j]

			if front == 0 {
				front = 1
			} else {
				front = 0
			}

			if back == 0 {
				back = 1
			} else {
				back = 0
			}

			image[i][j], image[i][len(image[i])-1-j] = back, front
		}
	}

	return image
}
