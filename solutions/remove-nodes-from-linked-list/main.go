package main

import "fmt"

func main() {
	node := listToLinkedList([]int{5, 2, 13, 3, 8})
	res := removeNodes(node)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1, 1, 1, 1})
	res = removeNodes(node)
	fmt.Println(linkedListToList(res))
}

func listToLinkedList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < n; i++ {
		node := &ListNode{Val: nums[i]}
		cur.Next = node
		cur = node
	}

	return head
}

func linkedListToList(head *ListNode) []int {
	res := make([]int, 0)

	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	next := removeNodes(head.Next)
	if next.Val > head.Val {
		return next
	}

	head.Next = next
	return head
}
