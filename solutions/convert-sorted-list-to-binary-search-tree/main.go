package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	sortedNums := linkedListToSlice(head)
	insertOrders := getInsertOrder(sortedNums)

	var root *TreeNode
	for _, num := range insertOrders {
		root = insertBST(root, num)
	}

	return root
}

func insertBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	} else {
		root.Right = insertBST(root.Right, val)
	}

	return root
}

func getInsertOrder(sortedNums []int) []int {
	n := len(sortedNums)

	if n == 0 {
		return []int{}
	}

	mid := n / 2
	midNum := sortedNums[mid]

	results := []int{midNum}

	leftSortedNums := sortedNums[:mid]
	rightSortedNums := sortedNums[mid+1:]

	leftResults := getInsertOrder(leftSortedNums)
	rightResults := getInsertOrder(rightSortedNums)

	results = append(results, leftResults...)
	results = append(results, rightResults...)

	return results
}

func linkedListToSlice(head *ListNode) []int {
	results := make([]int, 0)

	cur := head

	for cur != nil {
		results = append(results, cur.Val)
		cur = cur.Next
	}

	return results
}
