package main

import "fmt"

func main() {
	node := listToLinkedList([]int{1, 2, 3, 4})
	res := swapPairs(node)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{})
	res = swapPairs(node)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1})
	res = swapPairs(node)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1, 2, 3, 4, 5})
	res = swapPairs(node)
	fmt.Println(linkedListToList(res))

	node = listToLinkedList([]int{1, 2})
	res = swapPairs(node)
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

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	setHead := true
	var prev *ListNode

	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = next.Next
		next.Next = cur

		if setHead {
			head = next
			setHead = false
		} else {
			prev.Next = next
		}

		prev = cur
		cur = cur.Next
	}

	return head
}
