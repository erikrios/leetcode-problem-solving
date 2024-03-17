package main

import "fmt"

func main() {
	mergedNode := mergeInBetween(listToLinkedList([]int{0, 1, 2, 3, 4, 5}), 3, 4, listToLinkedList([]int{1000000, 1000001, 1000002}))
	fmt.Println(linkedListToList(mergedNode))

	mergedNode = mergeInBetween(listToLinkedList([]int{0, 1, 2, 3, 4, 5, 6}), 2, 5, listToLinkedList([]int{1000000, 1000001, 1000002, 1000003, 1000004}))
	fmt.Println(linkedListToList(mergedNode))
}

func listToLinkedList(list []int) *ListNode {
	if list == nil {
		return nil
	}

	head := &ListNode{Val: list[0]}
	cur := head

	for i := 1; i < len(list); i++ {
		node := &ListNode{Val: list[i]}
		cur.Next = node
		cur = node
	}

	return head
}

func linkedListToList(node *ListNode) []int {
	res := make([]int, 0)

	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var prevA, nextB *ListNode

	cur := list1
	var i int
	for cur != nil {
		if i == a-1 {
			prevA = cur
		}
		if i == b+1 {
			nextB = cur
			break
		}
		cur = cur.Next
		i++
	}

	cur = list2
	for cur.Next != nil {
		cur = cur.Next
	}

	prevA.Next = list2
	cur.Next = nextB

	return list1
}
