package main

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9))
	fmt.Println(shiftGrid([][]int{{1}}, 100))
	fmt.Println(shiftGrid([][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}}, 23))
}

func shiftGrid(grid [][]int, k int) [][]int {
	ll := NewCircularLinkedList()

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			val := grid[i][j]
			ll.Insert(val)
		}
	}

	ll.Shift(k)

	currentNode := ll.Head
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = currentNode.Val
			currentNode = currentNode.Next
		}
	}

	return grid
}

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type CircularLinkedList struct {
	Head *Node
	Tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{}
}

func (c *CircularLinkedList) Insert(v int) {
	n := &Node{Val: v}

	if c.Head == nil {
		n.Next = n
		n.Prev = n
		c.Head = n
		c.Tail = n
		return
	}

	n.Prev = c.Tail
	n.Next = c.Head
	c.Tail.Next = n
	c.Tail = n
	c.Head.Prev = c.Tail
}

func (c *CircularLinkedList) Shift(times int) {
	for i := 0; i < times; i++ {
		c.Head = c.Tail
		c.Tail = c.Tail.Prev
	}
}
