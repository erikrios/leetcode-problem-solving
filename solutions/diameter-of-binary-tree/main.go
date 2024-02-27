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
		},
	}
	fmt.Println(diameterOfBinaryTree(node))

	node = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	fmt.Println(diameterOfBinaryTree(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxD int

func diameterOfBinaryTree(root *TreeNode) int {
	maxD = 0
	calculateDiameter(root)
	return maxD
}

func calculateDiameter(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftVal := calculateDiameter(root.Left)
	rightVal := calculateDiameter(root.Right)

	maxLocal := leftVal + rightVal
	maxD = max(maxLocal, maxD)

	return max(leftVal, rightVal) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
