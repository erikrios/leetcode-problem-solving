package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(countNodes(node))

	fmt.Println(countNodes(nil))

	node = &TreeNode{Val: 1}
	fmt.Println(countNodes(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if isLeaf(root) {
		return 1
	}

	var count int
	count++
	if root.Left != nil {
		count += countNodes(root.Left)
	}
	if root.Right != nil {
		count += countNodes(root.Right)
	}

	return count
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
