package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 7))
	fmt.Println(busyStudent([]int{4}, []int{4}, 4))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var counter int

	for i := 0; i < len(startTime); i++ {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			counter++
		}
	}

	return counter
}
