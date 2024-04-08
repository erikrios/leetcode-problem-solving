package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(checkValidString("()"))
	fmt.Println(checkValidString("(*)"))
	fmt.Println(checkValidString("(*))"))
	fmt.Println(checkValidString("(*)))"))
	fmt.Println(checkValidString("((((()(()()()*()(((((*)()*(**(())))))(())()())(((())())())))))))(((((())*)))()))(()((*()*(*)))(*)()"))
	fmt.Println(checkValidString("("))
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"))
}

func checkValidString(s string) bool {
	stack := NewStack()
	asterixStack := NewStack()

	for i := 0; i < len(s); i++ {
		if char := s[i]; char == '(' {
			stack.Push(i)
		} else if char == '*' {
			asterixStack.Push(i)
		} else if char == ')' {
			if !stack.IsEmpty() {
				_ = stack.Pop()
			} else if !asterixStack.IsEmpty() {
				_ = asterixStack.Pop()
			} else {
				return false
			}
		}
	}

	for !stack.IsEmpty() && !asterixStack.IsEmpty() {
		stackTop, _ := stack.Peek()
		asterixTop, _ := asterixStack.Peek()
		if stackTop > asterixTop {
			return false
		}
		_ = stack.Pop()
		_ = asterixStack.Pop()
	}

	return stack.IsEmpty()
}

type stack struct {
	chars []int
}

func NewStack() *stack {
	return &stack{chars: make([]int, 0)}
}

func (s *stack) Push(c int) {
	s.chars = append(s.chars, c)
}

func (s *stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("empty stack")
	}

	s.chars = s.chars[:len(s.chars)-1]
	return nil
}

func (s *stack) Peek() (int, error) {
	if s.IsEmpty() {
		return ' ', errors.New("empty stack")
	}

	return s.chars[len(s.chars)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.chars) == 0
}
