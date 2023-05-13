package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}

type Distances [][]int

func (d Distances) Len() int {
	return len(d)
}

func (d Distances) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d Distances) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func allCellsDistOrder(rows, cols, rCenter, cCenter int) [][]int {
	var distances Distances = make([][]int, 0, rows)

	for i := 0; i < rows; i++ {
		var distR int
		if i < rCenter {
			distR = rCenter - i
		} else {
			distR = i - rCenter
		}

		for j := 0; j < cols; j++ {
			var distC int
			if j < cCenter {
				distC = cCenter - j
			} else {
				distC = j - cCenter
			}

			distances = append(distances, []int{distR + distC, i, j})
		}
	}

	sort.Sort(distances)

	for i := 0; i < len(distances); i++ {
		distances[i] = distances[i][1:]
	}

	return distances
}
