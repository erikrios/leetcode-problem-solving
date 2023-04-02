package main

import "fmt"

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}

	fmt.Println(sumRootToLeaf(node))

	node = &TreeNode{Val: 0}

	fmt.Println(sumRootToLeaf(node))

	node = &TreeNode{Val: 1, Left: &TreeNode{Val: 1}}

	fmt.Println(sumRootToLeaf(node))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	sum := sum(root, 0)
	return sum
}

func sum(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	val = val*2 + root.Val

	if root.Left == root.Right {
		return val
	} else {
		return sum(root.Left, val) + sum(root.Right, val)
	}
}
