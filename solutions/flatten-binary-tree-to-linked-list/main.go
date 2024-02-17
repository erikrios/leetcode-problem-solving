package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	flatten(node)
	PrintTree(node)

	node = nil
	flatten(node)
	PrintTree(node)

	node = &TreeNode{Val: 0}
	flatten(node)
	PrintTree(node)
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PrintTree(root.Left)
	PrintTree(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flattenToLinkedList(root)
}

func flattenToLinkedList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if isLeaf(root) {
		return root
	}

	left := flattenToLinkedList(root.Left)
	right := flattenToLinkedList(root.Right)

	if left != nil {
		left.Right = root.Right
		root.Right = root.Left
	}

	root.Left = nil

	if right != nil {
		return right
	} else if left != nil {
		return left
	} else {
		return root
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
