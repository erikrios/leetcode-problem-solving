package main

import "fmt"

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 2))
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{5, 1, 4, 2, 2}, 6))
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	var counter int

	for i := 0; i < len(hours); i++ {
		if hours[i] >= target {
			counter++
		}
	}

	return counter
}
