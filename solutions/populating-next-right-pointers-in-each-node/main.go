package main

func main() {
	node := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	connect(node)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfs(root, NewStack())
	return root
}

func dfs(root *Node, stack *Stack) {
	if !stack.IsEmpty() {
		root.Next = stack.Pop()
	}

	if isLeaf(root) {
		stack.Push(root)
		return
	}

	dfs(root.Right, stack)
	dfs(root.Left, stack)
	stack.Push(root)
}

func isLeaf(root *Node) bool {
	return root.Left == nil && root.Right == nil
}

type MyListNode struct {
	Val  *Node
	Next *MyListNode
}

type Stack struct {
	head *MyListNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val *Node) {
	node := &MyListNode{Val: val}
	node.Next = s.head
	s.head = node
}

func (s *Stack) Pop() *Node {
	if s.IsEmpty() {
		return nil
	}

	val := s.head.Val
	s.head = s.head.Next
	return val
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
