package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperatures([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperatures([]int{30, 60, 90}))
}

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := make([]struct {
		i int
		v int
	}, 0)
	answers := make([]int, n)

	for i := 0; i < n; i++ {
		temp := temperatures[i]

		for len(stack) > 0 {
			m := len(stack)
			prevTemp := stack[m-1]
			if prevTemp.v < temp {
				answers[prevTemp.i] = i - prevTemp.i
				stack = stack[:m-1]
			} else {
				break
			}
		}

		stack = append(stack, struct {
			i int
			v int
		}{i, temp})
	}

	return answers
}
