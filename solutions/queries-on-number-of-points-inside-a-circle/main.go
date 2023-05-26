package main

import "fmt"

func main() {
	fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
	fmt.Println(countPoints([][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}}, [][]int{{1, 2, 2}, {2, 2, 2}, {4, 3, 2}, {4, 3, 3}}))
}

// Point is a x, y representation of Vector 2D
type Point struct {
	X, Y int
}

// Circle represents the 2D circle
type Circle struct {
	Radius int
	Center Point
}

func (c *Circle) IsInside(p Point) bool {
	v := (p.X-c.Center.X)*(p.X-c.Center.X) + (p.Y-c.Center.Y)*(p.Y-c.Center.Y)
	r2 := c.Radius * c.Radius

	return v <= r2
}

func countPoints(points [][]int, queries [][]int) []int {
	results := make([]int, 0, len(queries))
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		circle := Circle{Radius: query[2], Center: Point{X: query[0], Y: query[1]}}

		var counter int
		for j := 0; j < len(points); j++ {
			p := points[j]
			x, y := p[0], p[1]
			point := Point{X: x, Y: y}
			if circle.IsInside(point) {
				counter++
			}
		}

		results = append(results, counter)
	}

	return results
}
