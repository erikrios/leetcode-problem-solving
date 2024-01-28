package main

import "fmt"

func main() {
	node := listToLinkedList([]int{1, 2, 6, 3, 4, 5, 6})
	res := removeElements(node, 6)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{})
	res = removeElements(node, 1)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{7, 7, 7, 7})
	res = removeElements(node, 7)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1, 2, 2, 1})
	res = removeElements(node, 2)
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

func removeElements(head *ListNode, val int) *ListNode {
	cur := head
	var prev *ListNode

	for cur != nil {
		if cur.Val == val {
			if cur == head {
				head = cur.Next
			} else {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}
		cur = cur.Next
	}

	return head
}
