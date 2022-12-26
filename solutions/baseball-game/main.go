package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	fmt.Println(calPoints([]string{"1", "C"}))
}

func calPoints(operations []string) int {
	stack := New()

	for _, operation := range operations {
		switch operation {
		case "+":
			a, _ := stack.Pop()
			b, _ := stack.Peek()
			stack.Push(a)
			stack.Push(a + b)
		case "D":
			prev, _ := stack.Peek()
			stack.Push(prev * 2)
		case "C":
			stack.Pop()
		default:
			intVal, _ := strconv.Atoi(operation)
			stack.Push(intVal)
		}
	}

	var res int
	for !stack.isEmpty() {
		v, _ := stack.Pop()
		res += v
	}

	return res
}

var (
	StackEmptyErr = errors.New("stack is empty")
)

type Stack struct {
	items []int
}

func New() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (v int, err error) {
	if s.isEmpty() {
		err = StackEmptyErr
		return
	}

	v = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func (s *Stack) Peek() (v int, err error) {
	if s.isEmpty() {
		err = StackEmptyErr
		return
	}

	v = s.items[len(s.items)-1]
	return
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
