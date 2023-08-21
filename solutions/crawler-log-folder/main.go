package main

import "fmt"

func main() {
	fmt.Println(minOperations([]string{"d1/", "d2/", "../", "d21/", "./"}))
	fmt.Println(minOperations([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}))
	fmt.Println(minOperations([]string{"d1/", "../", "../", "../"}))
}

func minOperations(logs []string) int {
	ll := NewLinkedList()
	ll.Insert("/")

	for i := 0; i < len(logs); i++ {
		switch log := logs[i]; log {
		case "../":
			if ll.Tail.Val != "/" {
				ll.Remove()
			}
		case "./":
			// Do nothing
		default:
			ll.Insert(log)
		}
	}

	return ll.Size() - 1
}

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head  *Node
	Tail  *Node
	Count int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Insert(val string) {
	node := &Node{Val: val}
	l.Count++

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node
}

func (l *LinkedList) Remove() {
	if l.Head == nil {
		return
	}

	l.Tail = l.Tail.Prev
	l.Tail.Next = nil

	l.Count--
}

func (l *LinkedList) Size() int {
	return l.Count
}
