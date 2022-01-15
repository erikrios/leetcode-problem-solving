package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	add(l1, l2)
	return l1
}

func add(l1, l2 *ListNode) *ListNode {
	if l1.Next == nil {
		l1.Next = &ListNode{}
	}

	if l2 != nil {
		l1.Val += l2.Val
		l2 = l2.Next
	}

	if l1.Val > 9 {
		l1.Val -= 10
		l1.Next.Val += 1
	}

	if l1.Next.Val == 0 && l1.Next.Next == nil && l2 == nil {
		l1.Next = nil
		return nil
	}

	return add(l1.Next, l2)
}
