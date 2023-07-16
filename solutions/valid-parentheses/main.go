package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

var openingBraces = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

var closingBraces = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	stack := &stack{}

	for i := 0; i < len(s); i++ {
		brace := s[i]

		if isOpeningBrace(brace) {
			stack.Push(brace)
		} else {
			openingPair := closingBraces[brace]
			v, err := stack.Peek()
			if err != nil || v != openingPair {
				return false
			}

			stack.Pop()
		}
	}

	return stack.IsEmpty()
}

func isOpeningBrace(brace byte) bool {
	_, ok := openingBraces[brace]
	return ok
}

var stackEmptyErr = errors.New("stack is empty")

type stack struct {
	braces []byte
}

func (s *stack) Pop() error {
	if s.IsEmpty() {
		return stackEmptyErr
	}

	s.braces = s.braces[0 : len(s.braces)-1]
	return nil
}

func (s *stack) Push(brace byte) {
	s.braces = append(s.braces, brace)
}

func (s *stack) Peek() (byte, error) {
	if s.IsEmpty() {
		return ' ', stackEmptyErr
	}

	return s.braces[len(s.braces)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.braces) == 0
}
