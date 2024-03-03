package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
}

func minAddToMakeValid(s string) int {
	stack := NewStack()

	for i := 0; i < len(s); i++ {
		switch char := s[i]; char {
		case '(':
			stack.Push(char)
		default:
			topChar := stack.Peek()
			if topChar == ')' || stack.IsEmpty() {
				stack.Push(char)
			} else {
				stack.Pop()
			}
		}
	}

	return stack.Size()
}

type Stack struct {
	head *SinglyListNode
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val byte) {
	node := &SinglyListNode{Val: val, Next: s.head}
	s.head = node
	s.size++
}

func (s *Stack) Peek() byte {
	if s.head == nil {
		return ' '
	}
	return s.head.Val
}

func (s *Stack) Pop() byte {
	if s.head == nil {
		return ' '
	}

	val := s.head.Val
	s.head = s.head.Next
	s.size--

	return val
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

type SinglyListNode struct {
	Val  byte
	Next *SinglyListNode
}
