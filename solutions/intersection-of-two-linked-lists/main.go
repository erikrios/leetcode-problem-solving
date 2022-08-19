package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapper := make(map[*ListNode]bool)

	for current := headA; current != nil; current = current.Next {
		mapper[current] = !mapper[current]
	}

	for current := headB; current != nil; current = current.Next {
		mapper[current] = !mapper[current]
		if !mapper[current] {
			break
		}
	}

	for k, v := range mapper {
		if !v {
			return k
		}
	}

	return nil
}
