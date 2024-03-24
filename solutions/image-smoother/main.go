package main

import "math"

func main() {
}

func imageSmoother(img [][]int) [][]int {
	m := len(img)
	n := len(img[0])

	results := make([][]int, m)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			results[i] = append(results[i], surroundingAvg(img, i, j))
		}
	}

	return results
}

func surroundingAvg(img [][]int, centerRowIndex, centerColIndex int) int {
	m := len(img)
	n := len(img[0])

	var sum, divider int

	center := img[centerRowIndex][centerColIndex]
	sum += center
	divider++

	if centerRowIndex-1 >= 0 && centerColIndex-1 >= 0 {
		topLeft := img[centerRowIndex-1][centerColIndex-1]
		sum += topLeft
		divider++
	}

	if centerRowIndex-1 >= 0 {
		top := img[centerRowIndex-1][centerColIndex]
		sum += top
		divider++
	}

	if centerRowIndex-1 >= 0 && centerColIndex+1 < n {
		topRight := img[centerRowIndex-1][centerColIndex+1]
		sum += topRight
		divider++
	}

	if centerColIndex-1 >= 0 {
		centerLeft := img[centerRowIndex][centerColIndex-1]
		sum += centerLeft
		divider++
	}

	if centerColIndex+1 < n {
		centerRight := img[centerRowIndex][centerColIndex+1]
		sum += centerRight
		divider++
	}

	if centerRowIndex+1 < m && centerColIndex-1 >= 0 {
		bottomLeft := img[centerRowIndex+1][centerColIndex-1]
		sum += bottomLeft
		divider++
	}

	if centerRowIndex+1 < m {
		bottom := img[centerRowIndex+1][centerColIndex]
		sum += bottom
		divider++
	}

	if centerRowIndex+1 < m && centerColIndex+1 < n {
		bottomRight := img[centerRowIndex+1][centerColIndex+1]
		sum += bottomRight
		divider++
	}

	return int(math.Floor(float64(sum) / float64(divider)))
}
