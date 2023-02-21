package main

import "fmt"

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{{1, 1}, {3, 4}, {-1, 0}}))
	fmt.Println(minTimeToVisitAllPoints([][]int{{3, 2}, {-2, 2}}))
}

func minTimeToVisitAllPoints(points [][]int) int {
	current := points[0]
	var total int
	for i := 1; i < len(points); i++ {
		point := points[i]

		var counter int
		for {
			if current[0] < point[0] {
				current[0]++
			} else if current[0] > point[0] {
				current[0]--
			}

			if current[1] < point[1] {
				current[1]++
			} else if current[1] > point[1] {
				current[1]--
			}

			counter++
			if current[0] == point[0] && current[1] == point[1] {
				break
			}
		}

		total += counter
	}

	return total
}
