package main

import "fmt"

func main() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
	fmt.Println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var counter int

	for i := 0; i <= len(arr)-3; i++ {
		for j := i + 1; j <= len(arr)-2; j++ {
			for k := j + 1; k <= len(arr)-1; k++ {
				valA, valB, valC := arr[i]-arr[j], arr[j]-arr[k], arr[i]-arr[k]

				if valA < 0 {
					valA *= -1
				}
				if valB < 0 {
					valB *= -1
				}
				if valC < 0 {
					valC *= -1
				}

				if valA <= a && valB <= b && valC <= c {
					counter++
				}
			}
		}
	}

	return counter
}
