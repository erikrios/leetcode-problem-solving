package main

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
	fmt.Println(replaceElements([]int{400}))
}

func replaceElements(arr []int) []int {
	results := make([]int, len(arr), len(arr))

	max := arr[len(arr)-1]
	for i := len(arr) - 1 - 1; i >= 0; i-- {
		results[i] = max
		if num := arr[i]; num > max {
			max = num
		}
	}

	results[len(arr)-1] = -1

	return results
}
