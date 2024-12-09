package main

import "fmt"

func main() {
	node1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	node2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
	}
	fmt.Println(flipEquiv(node1, node2))

	node1 = nil
	node2 = nil
	fmt.Println(flipEquiv(node1, node2))

	node1 = nil
	node2 = &TreeNode{Val: 1}
	fmt.Println(flipEquiv(node1, node2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil && root2 != nil {
		return false
	} else if root1 != nil && root2 == nil {
		return false
	} else if root1.Val != root2.Val {
		return false
	}

	if isLeaf(root1) && isLeaf(root2) {
		return true
	}

	if root1.Left != nil && root1.Right != nil && root2.Left != nil && root2.Right != nil {
		if root1.Left.Val == root2.Left.Val && root1.Right.Val == root2.Right.Val {
			return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
		} else if root1.Left.Val == root2.Right.Val && root1.Right.Val == root2.Left.Val {
			swapChildren(root1)
			return flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)
		} else {
			return false
		}
	} else if root1.Left != nil && root1.Right == nil && root2.Left != nil && root2.Right == nil {
		if root1.Left.Val != root2.Left.Val {
			return false
		} else {
			return flipEquiv(root1.Left, root2.Left)
		}
	} else if root1.Left == nil && root1.Right != nil && root2.Left == nil && root2.Right != nil {
		if root1.Right.Val != root2.Right.Val {
			return false
		} else {
			return flipEquiv(root1.Right, root2.Right)
		}
	} else if root1.Left != nil && root1.Right == nil && root2.Left == nil && root2.Right != nil {
		if root1.Left.Val != root2.Right.Val {
			return false
		} else {
			swapChildren(root1)
			return flipEquiv(root1.Right, root2.Right)
		}
	} else if root1.Left == nil && root1.Right != nil && root2.Left != nil && root2.Right == nil {
		if root1.Right.Val != root2.Left.Val {
			return false
		} else {
			swapChildren(root1)
			return flipEquiv(root1.Left, root2.Left)
		}
	} else {
		return false
	}
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func swapChildren(root *TreeNode) {
	root.Left, root.Right = root.Right, root.Left
}
