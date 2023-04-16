package main

import (
	"fmt"
)

func main() {
	node1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	node2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}

	res := mergeTrees(node1, node2)
	printTree(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	node := &TreeNode{Val: root1.Val + root2.Val}
	node.Left = mergeTrees(root1.Left, root2.Left)
	node.Right = mergeTrees(root1.Right, root2.Right)
	return node
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
