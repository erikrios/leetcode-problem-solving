package main

import "fmt"

func main() {
	node := listToLinkedList([]int{1, 2, 3})
	results := splitListToParts(node, 5)
	for _, v := range results {
		fmt.Println(linkedListToList(v))
	}

	node = listToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	results = splitListToParts(node, 3)
	for _, v := range results {
		fmt.Println(linkedListToList(v))
	}

	node = listToLinkedList([]int{1})
	results = splitListToParts(node, 3)
	for _, v := range results {
		fmt.Println(linkedListToList(v))
	}

	results = splitListToParts(nil, 1)
	for _, v := range results {
		fmt.Println(linkedListToList(v))
	}

	node = listToLinkedList([]int{0, 1, 2})
	results = splitListToParts(node, 2)
	for _, v := range results {
		fmt.Println(linkedListToList(v))
	}
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

func splitListToParts(head *ListNode, k int) []*ListNode {
	results := make([]*ListNode, k)

	if head == nil {
		return results
	}

	slow, fast := head, head

	i := 1
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		i++
	}

	var n int
	if fast.Next == nil { // odd
		n = (i * 2) - 1
	} else { // even
		n = i * 2
	}

	var mod int
	var div int
	if n > k {
		mod = n % k
		div = n / k
	} else {
		div = 1
	}

	cur := head
	var newHead *ListNode
	i = 0
	divCounter := div
	var isModAlreadyAdded bool
	shouldAdd := true
	for cur != nil {
		if shouldAdd {
			newHead = cur
			results[i] = newHead
			shouldAdd = false
		}
		prev := cur
		cur = cur.Next
		divCounter--

		if divCounter == 0 {
			if mod > 0 && !isModAlreadyAdded {
				divCounter++
				mod--
				isModAlreadyAdded = true
			} else {
				i++
				isModAlreadyAdded = false
				divCounter = div
				prev.Next = nil
				shouldAdd = true
			}
		}
	}

	return results
}
