package main

import "fmt"

func main() {
	fmt.Println(sumSubarrayMins([]int{3, 1, 2, 4}))
	fmt.Println(sumSubarrayMins([]int{11, 81, 94, 43, 3}))
}

func sumSubarrayMins(arr []int) int {
	mod := 1_000_000_000 + 7
	n := len(arr)
	stack := make([][2]int, 0)
	var sum int

	for i := 0; i < n; i++ {
		val := arr[i]

		for len(stack) > 0 && val < stack[len(stack)-1][1] {
			curStack := stack[len(stack)-1]
			curStackVal := curStack[1]
			curStackIdx := curStack[0]

			// popping from the stack
			stack = stack[:len(stack)-1]

			var left, right int

			if len(stack) > 0 {
				left = curStackIdx - stack[len(stack)-1][0]
			} else {
				left = curStackIdx + 1
			}

			right = i - curStackIdx

			res := curStackVal * left * right
			sum = (sum + res) % mod
		}

		// pushing to the stack
		stack = append(stack, [2]int{i, val})
	}

	for i := 0; i < len(stack); i++ {
		curStack := stack[i]
		curStackVal := curStack[1]
		curStackIdx := curStack[0]
		var left, right int

		if i > 0 {
			left = curStackIdx - stack[i-1][0]
		} else {
			left = curStackIdx + 1
		}

		right = len(arr) - curStackIdx

		res := curStackVal * left * right
		sum = (sum + res) % mod
	}

	return sum
}
