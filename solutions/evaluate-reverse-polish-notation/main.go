package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(token []string) int {
	stack := make([]int, 0)

	for i := 0; i < len(token); i++ {
		val := token[i]
		if num, err := strconv.Atoi(val); err != nil {
			n := len(stack)
			num2 := stack[n-1]
			num1 := stack[n-2]
			stack = stack[:n-2]
			var res int
			switch char := val[0]; char {
			case '+':
				res = num1 + num2
			case '-':
				res = num1 - num2
			case '*':
				res = num1 * num2
			default:
				res = num1 / num2
			}
			stack = append(stack, res)
		} else {
			stack = append(stack, num)
		}
	}

	return stack[0]
}
