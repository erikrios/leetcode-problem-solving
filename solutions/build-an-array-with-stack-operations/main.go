package main

import "fmt"

func main() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
	fmt.Println(buildArray([]int{2, 3, 4, 5, 8, 9, 10}, 10))
}

func buildArray(target []int, n int) []string {
	operations := make([]string, 0)
	var streamVal int
	var defisit int

	for i := 0; i < len(target); {
		num := target[i]
		streamVal++
		n--
		if streamVal == num {
			if defisit == 0 {
				operations = append(operations, "Push")
			} else {
				for defisit > 0 {
					operations = append(operations, "Push", "Pop")
					defisit--
				}
				operations = append(operations, "Push")
			}
			i++
		} else {
			defisit++
		}
	}

	return operations
}
