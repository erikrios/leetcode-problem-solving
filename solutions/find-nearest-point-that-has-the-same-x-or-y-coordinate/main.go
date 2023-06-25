package main

import "fmt"

func main() {
	fmt.Println(nearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(nearestValidPoint(3, 4, [][]int{{2, 3}}))
}

func nearestValidPoint(x, y int, points [][]int) int {
	index := -1
	manhattanDist := 10_000

	for i := 0; i < len(points); i++ {
		point := points[i]

		if point[0] == x || point[1] == y {
			var absX, absY int

			if x > point[0] {
				absX = x - point[0]
			} else {
				absX = point[0] - x
			}

			if y > point[1] {
				absY = y - point[1]
			} else {
				absY = point[1] - y
			}

			dist := absX + absY
			if index == -1 {
				index = i
				manhattanDist = dist
				continue
			}

			if dist < manhattanDist {
				index = i
				manhattanDist = dist
			}
		}
	}

	return index
}
