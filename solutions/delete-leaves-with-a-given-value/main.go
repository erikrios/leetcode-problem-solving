package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	node = removeLeafNodes(node, 2)

	printNodes(node)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printNodes(root *TreeNode) {
	if root == nil {
		return
	}

	printNodes(root.Left)
	fmt.Println(root.Val)
	printNodes(root.Right)
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if remove(root, target) && root.Val == target {
		return nil
	}

	return root
}

func remove(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}

	if isLeaf(root) {
		if root.Val == target {
			return true
		}

		return false
	}

	isLeftRemoved := remove(root.Left, target)
	if isLeftRemoved {
		root.Left = nil
	}

	isRightRemoved := remove(root.Right, target)
	if isRightRemoved {
		root.Right = nil
	}

	if isLeftRemoved && isRightRemoved && root.Val == target {
		return true
	}

	return false
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
