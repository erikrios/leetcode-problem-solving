package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(maxDepth(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMaxDepth := maxDepth(root.Left)
	rightMaxDepth := maxDepth(root.Right)

	var max int
	if leftMaxDepth > rightMaxDepth {
		max = leftMaxDepth
	} else {
		max = rightMaxDepth
	}

	return max + 1
}
