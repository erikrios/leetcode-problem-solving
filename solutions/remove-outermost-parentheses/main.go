package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}

func removeOuterParentheses(s string) string {
	stack := New()
	var emptyState bool
	results := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]

		if stack.isEmpty() && char == '(' {
			emptyState = true
		} else {
			emptyState = false
		}

		if char == '(' {
			stack.Push(char)
		} else {
			stack.Pop()
		}

		if emptyState {
			continue
		}

		if stack.isEmpty() && char == ')' {
			emptyState = true
		} else {
			emptyState = false
		}

		if emptyState {
			continue
		}

		results = append(results, char)
	}

	return string(results)
}

type Stack struct {
	items []byte
}

func New() *Stack {
	return &Stack{items: make([]byte, 0)}
}

func (s *Stack) Push(v byte) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (v byte) {
	v = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
