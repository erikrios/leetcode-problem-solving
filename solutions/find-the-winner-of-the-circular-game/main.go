package main

import "fmt"

func main() {
	fmt.Println(findTheWinner(5, 2))
	fmt.Println(findTheWinner(6, 5))
}

func findTheWinner(n, k int) int {
	head := &MyListNode{Val: 1}
	cur := head

	for i := 2; i <= n; i++ {
		cur.Next = &MyListNode{Val: i}
		cur = cur.Next
	}

	cur.Next = head

	prev := cur
	for head.Next != head {
		for i := 0; i < k-1; i++ {
			prev = head
			head = head.Next
		}

		prev.Next = head.Next
		head = head.Next
	}

	return head.Val
}

type MyListNode struct {
	Val  int
	Next *MyListNode
}
