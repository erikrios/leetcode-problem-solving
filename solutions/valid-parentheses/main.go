package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	openingBraces := []byte{'(', '{', '['}
	closingBraces := []byte{')', '}', ']'}

	stack := &stack{}

	for i := 0; i < len(s); i++ {
		if s[i] == openingBraces[0] || s[i] == openingBraces[1] || s[i] == openingBraces[2] {
			stack.Push(s[i])
		} else if (s[i] == closingBraces[0] && !stack.IsEmpty() && stack.Peek() == openingBraces[0]) || (s[i] == closingBraces[1] && !stack.IsEmpty() && stack.Peek() == openingBraces[1]) || (s[i] == closingBraces[2] && !stack.IsEmpty() && stack.Peek() == openingBraces[2]) {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}

	return stack.IsEmpty()
}

type stack struct {
	braces []byte
}

func (s *stack) Pop() {
	s.braces = s.braces[0 : len(s.braces)-1]
}

func (s *stack) Push(brace byte) {
	s.braces = append(s.braces, brace)
}

func (s *stack) Peek() byte {
	return s.braces[len(s.braces)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.braces) == 0
}
